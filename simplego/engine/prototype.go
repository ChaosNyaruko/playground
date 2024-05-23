package engine

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"

	"github.com/traefik/yaegi/interp"
)

var c1 = `
	import (
		. "fake.com/engine/proto"
	)

	func WrappedC1(bound int) func(*StreamInfo, *ClientInfo) (int, int, int, int) {
		x := bound
		return func(s *StreamInfo, c *ClientInfo) (id, reset, code, quit int) {
			if s.Concurrent > x {
				return 2003, 1, 233, 1
			}
			if c.Platform == 4 {
				return 2002, 0, 234, 1
			}
			return 2002, 0, 235, 0
		}
	}
`

var c2 = `
	import (
		. "fake.com/engine/proto"
	)

	func WrappedC2() func(*StreamInfo, *ClientInfo) (int, int, int, int) {
		return func(s *StreamInfo, c *ClientInfo) (id, reset, code, quit int) {
			if s.Concurrent == 137 {
				return 2003, 1, 300, 1
			}
			return 2002, 0, 235, 0
		}
	}
`

// MVP版本，每个Hook点要传的参数人工注册进ruleengine SDK
type StreamInfo struct {
	Concurrent int
	ClientType int
	AppID      int
}

type ClientInfo struct {
	Platform      int
	ServerGroupID int
}

type PMPeerInfo struct {
}

type PredictFileVV struct {
}

func CheckContent(content string) (*ast.File, error) {
	fset := token.NewFileSet()
	ast, err := parser.ParseFile(fset, "test", nil, parser.SkipObjectResolution)

	return ast, err
}

func updateFunc[T any](f *T, i *interp.Interpreter, name string) {
	v, err := i.Eval(name)
	if err == nil {
		*f = v.Interface().(T)
	}
}

func ParseComponent[T any, P any](wrapper T, sig P, content string, name string, args ...any) (P, error) {
	var err error
	var res T
	log.Printf("wrapper: %T, sig: %T, res: %T", wrapper, sig, res)
	var out P
	r := reflect.TypeOf(wrapper) // Why not res, T is literally an interface
	if r.Kind() != reflect.Func {
		return out, fmt.Errorf("the res must be a func type, but got %v", r.Kind())
	}
	i := interp.New(interp.Options{})
	if err = i.Use(interp.Exports{
		"fake.com/engine/proto/.": {
			"StreamInfo": reflect.ValueOf((*StreamInfo)(nil)),
			"ClientInfo": reflect.ValueOf((*ClientInfo)(nil)),
		},
	}); err != nil {
		return out, err
	}

	if _, err = i.Eval(content); err != nil {
		return out, err
	}
	updateFunc(&res, i, name)

	// TODO: check correctness
	in := []reflect.Value{}
	for i := 0; i < len(args); i++ {
		in = append(in, reflect.ValueOf(args[i]))
	}
	f := reflect.ValueOf(res).Call(in)
	out = f[0].Interface().(P)
	return out, nil
}

func ParseAll() error {
	// TODO:
	// 1. Read from Database: pipeline table
	// 2. Parse "content", get the components' idx/parameters/sequence
	// 3. Call ParseComponent to get a native ComponentForHookX
	// 4. Call Compose to get the whole executable "pipeline", which is effectively a FUNCTION, and maybe should be responsible for registering in the "router"(different according to the PSM?)

	type biz struct {
		t          string
		sid        int
		appid      int
		clienttype int
		content    string
		name       string
	}
	// TODO: read from DB or some kind of metadata
	bizs := []biz{
		{"normal", 0, 1128, 4, "1(bound=50)+2", "dispatcher_grouping"},
		{"normal", 0, 1128, 1, "2+1(bound=50)", "dispatcher_grouping"},
	}
	for _, b := range bizs {
		p, err := ParsePipeline(b.content, b.name)
		if err != nil {
			return err
		}
		cs := []ComponentForHook1{}
		for _, rc := range p.content {
			c, _ := ParseComponent(rc.wrapper, hook1Stub, rc.content, rc.name, rc.params...)
			cs = append(cs, c)
		}
		compiled := Compose(cs)
		H.router.Register(b.t, b.sid, b.appid, b.clienttype, compiled)
	}
	return nil
}

type RawPipeline struct {
	// TODO
	// components idx and parameters
	// sequence
	// content
	content []contentWithParams
}

type contentWithParams struct {
	name    string
	content string
	params  []any
	wrapper interface{}
}

func ParsePipeline(content string, name string) (*RawPipeline, error) {
	// TODO: just a mock, need a real parser here, maybe a manual one rather than regex is easier
	if content == "1(bound=50)+2" {
		return &RawPipeline{
			content: []contentWithParams{{"WrappedC1", c1, []any{50}, wrapperC1Stub}, {"WrappedC2", c2, []any{}, wrapperC2Stub}},
		}, nil
	}
	return &RawPipeline{
		content: []contentWithParams{{"WrappedC2", c2, []any{}, wrapperC2Stub}, {"WrappedC1", c1, []any{50}, wrapperC1Stub}},
	}, nil
}

var ComponentForHookMap map[string]any = map[string]any{"ComponentForHook1": ComponentForHook1(nil)}

func OldParseComponent(content string, name string, args ...any) (ComponentForHook1, error) {
	var err error
	i := interp.New(interp.Options{})
	if err = i.Use(interp.Exports{
		"fake.com/engine/proto/.": {
			"StreamInfo": reflect.ValueOf((*StreamInfo)(nil)),
			"ClientInfo": reflect.ValueOf((*ClientInfo)(nil)),
		},
	}); err != nil {
		return nil, err
	}

	if _, err = i.Eval(content); err != nil {
		return nil, err
	}
	// TODO: better and more robust generics implementation?
	switch len(args) {
	case 0:
		var f WrappedComponent2ForHook1
		updateFunc(&f, i, name)
		return f(), err
	case 1:
		var f WrappedComponent1ForHook1
		updateFunc(&f, i, name)
		return f(args[0].(int)), err
	default:
		return nil, fmt.Errorf("wrong argument for Parsing components")
	}
}

func Compose(components []ComponentForHook1) PipelineForHook1 {
	return PipelineForHook1{
		Components: components,
	}
}

type PipelineForHook1 struct {
	Components []ComponentForHook1
}

func (p PipelineForHook1) Execute(s *StreamInfo, c *ClientInfo) (id, reset, code int) {
	quit := 0
	for _, checker := range p.Components {
		id, reset, code, quit = checker(s, c)
		if quit != 0 {
			return
		}
	}
	return
}

type Router struct {
	normal   map[int]map[int]PipelineForHook1
	override map[int]PipelineForHook1
}

func (r *Router) Init() error {
	r.normal = make(map[int]map[int]PipelineForHook1)
	r.override = make(map[int]PipelineForHook1)
	return nil
}

func (r *Router) Register(t string, sid, appID, clientType int, p PipelineForHook1) error {
	switch t {
	case "override":
		r.override[sid] = p
	case "normal":
		if _, ok := r.normal[appID]; !ok {
			r.normal[appID] = make(map[int]PipelineForHook1)
		}
		r.normal[appID][clientType] = p
	default:
		return fmt.Errorf("unexpected type for register: %v", t)
	}
	return nil
}

func (r *Router) Get(s *StreamInfo, c *ClientInfo) PipelineForHook1 {
	// TODO: just a demo, no error handling yet
	if p, ok := r.override[c.ServerGroupID]; ok {
		return p
	}
	return r.normal[s.AppID][s.ClientType]
}

type DispatcherHook1 struct {
	router Router
}

// 针对Hook1的两个Component（原子策略），外包装可以不一样，但返回的闭包签名是相同的，都是ComponentForHook1
type WrappedComponent1ForHook1 = func(int) ComponentForHook1
type WrappedComponent2ForHook1 = func() ComponentForHook1

// MVP版我们可以约定一定要有一个quit返回值
type ComponentForHook1 = func(s *StreamInfo, c *ClientInfo) (int, int, int, int)

var wrapperC1Stub = func(int) ComponentForHook1 {
	return func(s *StreamInfo, c *ClientInfo) (int, int, int, int) {
		return 0, 0, 0, 0
	}
}
var wrapperC2Stub = func() ComponentForHook1 {
	return func(s *StreamInfo, c *ClientInfo) (int, int, int, int) {
		return 0, 0, 0, 0
	}
}

var hook1Stub = func(*StreamInfo, *ClientInfo) (int, int, int, int) {
	return 0, 0, 0, 0
}

var H *DispatcherHook1 = &DispatcherHook1{}

func (h *DispatcherHook1) Init() {
	_ = h.router.Init()
}

func GetRuntimeGroupID(s *StreamInfo, c *ClientInfo) (int, int, int) {
	return H.router.Get(s, c).Execute(s, c)
}

func Init() error {
	H.Init()
	return ParseAll()
}
