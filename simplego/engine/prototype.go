package engine

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

var cany = `
	import (
		. "fake.com/engine/proto"
	)

	func WrappedAny(bound interface{}) func(*StreamInfo, *ClientInfo) (int, int, int, int) {
		x := bound.(int)
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

func updateFunc[T any](f *T, i *interp.Interpreter, name string) error {
	v, err := i.Eval(name)
	if err == nil {
		*f = v.Interface().(T)
	} else {
		panic(err)
	}
	return err
}

func ParseComponent[P any](wrapper any, sig P, content string, name string, args ...any) (P, error) {
	var err error
	var res any
	log.Printf("wrapper: %T, sig: %T, res: %T", wrapper, sig, res)
	var out P

	r := reflect.TypeOf(wrapper) // Why not res, T is literally an interface
	if r.Kind() != reflect.Func {
		return out, fmt.Errorf("the res must be a func type, but got %v", r.Kind())
	}
	i := interp.New(interp.Options{})
	if err = i.Use(stdlib.Symbols); err != nil {
		panic(err)
	}
	if err = i.Use(interp.Exports{
		"fake.com/engine/proto/.": {
			"StreamInfo": reflect.ValueOf((*StreamInfo)(nil)),
			"ClientInfo": reflect.ValueOf((*ClientInfo)(nil)),
			"Mixed":      reflect.ValueOf((*Mixed)(nil)),
		},
	}); err != nil {
		panic(err)
		return out, err
	}

	if _, err = i.Eval(content); err != nil {
		panic(err)
		return out, err
	}
	if err := updateFunc(&res, i, name); err != nil {
		return out, err
	}

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
		// {"normal", 0, 1128, 4, "1(bound=50)+2", "dispatcher_grouping"},
		// {"normal", 0, 1128, 1, "2+1(bound=50)", "dispatcher_grouping"},
		// {"normal", 0, 1128, 4, "WrappedC1(bound=50)+WrappedC2()", "dispatcher_grouping"},
		{"normal", 0, 1128, 4, "WrappedAny(bound=50)+WrappedC2()", "dispatcher_grouping"},
		{"normal", 0, 1128, 1, "WrappedC2()+WrappedC1(bound=50)", "dispatcher_grouping"},
		// {"normal", 0, 1128, 4, "WrappedC1(bound=50)+WrappedC2", "dispatcher_grouping"},
		// {"normal", 0, 1128, 1, "WrappedC3+WrappedC1(bound=50)", "dispatcher_grouping"},
	}
	for _, b := range bizs {
		p, err := ParseP(b.content, b.name)
		if err != nil {
			return err
		}
		cs := []ComponentForAny{}
		for _, rc := range p.content {
			c, err := ParseComponent(rc.wrapper, anyStub, rc.content, rc.name, rc.params...)
			if err != nil {
				log.Printf("content: %v", p.content)
				panic(err)
			}
			cs = append(cs, c)
		}
		compiled := ComposeAny(cs)
		A.router.Register(b.name, b.t, b.sid, b.appid, b.clienttype, compiled)
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

func ParseP(content string, name string) (*RawPipeline, error) {
	// Compile the regular expression
	re, err := regexp.Compile(`(\w+)\(((?:\s*\w+\s*=\s*[^,)]+\s*,?)*)\)`)
	if err != nil {
		return nil, fmt.Errorf("Error compiling regex: %v", err)
	}

	if re.MatchString(content) {
		fmt.Printf("Expression '%s' matches the regex\n", content)
	} else {
		return nil, fmt.Errorf("Expression '%s' does not match the regex\n", content)
	}

	type Parameter struct {
		Name  string
		Value string
	}
	type Expression struct {
		FunctionName string
		Parameters   []Parameter
	}
	res := &RawPipeline{
		content: []contentWithParams{},
	}
	var allExpressions []Expression
	fmt.Println(content, "------")
	subExpressions := strings.Split(content, "+")
	interpreter := interp.New(interp.Options{})
	for _, subExpr := range subExpressions {
		if matches := re.FindAllStringSubmatch(subExpr, -1); matches != nil {
			for _, match := range matches {
				var e Expression
				e.FunctionName = match[1]

				paramStr := match[2]
				params := strings.Split(paramStr, ",")
				for _, param := range params {
					if param = strings.TrimSpace(param); param != "" {
						parts := strings.Split(param, "=")
						if len(parts) == 2 {
							e.Parameters = append(e.Parameters, Parameter{
								Name:  strings.TrimSpace(parts[0]),
								Value: strings.TrimSpace(parts[1]),
							})
						}
					}
				}

				fmt.Printf("Expression: %s\n", e.FunctionName)
				// TODO: query DB to get the signature, i.e. the arguments' types
				// e.g. WrappedC1: -> "int"
				// e.g. WrappedC2: -> ""
				// e.g. WrappedC3: -> "string"
				// e.g. WrappedC4: -> "int,string,float64,bool"
				// this prototype are just mocking.
				types := map[string][]string{
					"WrappedC1":  []string{"int"},
					"WrappedC2":  []string{},
					"WrappedC3":  []string{"string"},
					"WrappedC4":  []string{"int", "string", "float64", "bool"},
					"WrappedAny": []string{"any"},
				}
				contents := map[string]string{
					"WrappedC1":  c1_,
					"WrappedC2":  c2_,
					"WrappedC3":  "",
					"WrappedC4":  "",
					"WrappedAny": cany_,
				}
				stubKey := fmt.Sprintf("<%s>", strings.Join(types[e.FunctionName], ","))
				stub := wrapperMap[stubKey]
				if stub == nil {
					fmt.Printf("nil stub for %q, please check!\n", e.FunctionName)
				}
				ele := contentWithParams{
					name:    e.FunctionName,
					content: contents[e.FunctionName],
					params:  []any{},
					wrapper: stub,
				}

				// TODO: we need a better validation system
				if len(types[e.FunctionName]) != len(e.Parameters) {
					return nil, fmt.Errorf("stubkey: %s, parameter number wrong for %v, registered: %d, but passed in %d",
						stubKey, e.FunctionName, len(types[e.FunctionName]), len(e.Parameters))
				}
				for i, param := range e.Parameters {
					fmt.Printf("  - %s = %s\n", param.Name, param.Value)
					t := types[e.FunctionName][i]
					tmp, err := interpreter.Eval(param.Value)
					fmt.Printf("-----=====%v:%v:%v\n", tmp, tmp.Interface(), tmp.Type())
					if err != nil || (types[e.FunctionName][i] != "any" && tmp.Type().String() != types[e.FunctionName][i]) {
						return nil, ErrWrongParameterType.WithArgs(i, param.Name, t, param.Value, tmp.Type().String())
					}
					ele.params = append(ele.params, tmp.Interface())
				}
				fmt.Println()
				res.content = append(res.content, ele)
			}
		} else {
			fmt.Printf("Expression '%s' does not match the regex\n", subExpr)
		}
	}

	for _, e := range allExpressions {
		fmt.Printf("Expression: %s\n", e.FunctionName)
		for _, param := range e.Parameters {
			fmt.Printf("  - %s = %s\n", param.Name, param.Value)
		}
		fmt.Println()
	}
	fmt.Printf("res: %+v, %v--\n", res, reflect.ValueOf(res))
	for i, x := range res.content {
		for j, y := range x.params {
			fmt.Printf("param[%d-%d], %v, %v, %v\n", i, j, y, reflect.TypeOf(y).Name(), reflect.ValueOf(y))
		}
	}
	fmt.Printf("--\n")
	return res, nil
}

// var wrapperMap = map[string]any{
// 	"<int>": wrapperC1Stub,
// 	"<>":    wrapperC2Stub,
// 	"<any>": wrapperAnyStub,
// }

var wrapperMap = map[string]any{
	"<int>": anyWrapperInt,
	"<>":    anyWrapperEmpty,
	"<any>": anyWrapperAny,
}

func ParsePipeline(content string, name string) (*RawPipeline, error) {
	// TODO: just a mock, need a real parser here, maybe a manual one rather than regex is easier
	if content == "WrappedC1(bound=50)+WrappedC2()" {
		return &RawPipeline{
			content: []contentWithParams{{"WrappedC1", c1, []any{50}, wrapperC1Stub}, {"WrappedC2", c2, []any{}, wrapperC2Stub}},
		}, nil
	}
	return &RawPipeline{
		content: []contentWithParams{{"WrappedC2", c2, []any{}, wrapperC2Stub}, {"WrappedC1", c1, []any{50}, wrapperC1Stub}},
	}, nil
	// ^(\w+)\((?:([\w\d]+)=(\d+(?:\.\d+)?))?,?(?:([\w\d]+)=(\d+(?:\.\d+)?))?\)$
	// ^(\w+)\((?:([\w\d]+)=([^,)]+))?,?(?:([\w\d]+)=([^,)]+))?\)$
	// ^((\w+)\((?:([\w\d]+)=([^,)]+))?,?(?:([\w\d]+)=([^,)]+))?\))(?:\+(\w+)\((?:([\w\d]+)=([^,)]+))?,?(?:([\w\d]+)=([^,)]+))?\))*$

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

var wrapperAnyStub = func(any) ComponentForHook1 {
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
	// H.Init()
	A.Init()
	return ParseAll()
}

func X() {
	type Parameter struct {
		Name  string
		Value string
	}
	type Expression struct {
		FunctionName string
		Parameters   []Parameter
	}

	// Compile the regular expression
	re, err := regexp.Compile(`(\w+)\((?:([\w\d]+)=([^,)]+))?(?:,\s*(?:([\w\d]+)=([^,)]+))?(?:,\s*(?:([\w\d]+)=([^,)]+))?)?)?\)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Test the expressions
	expressions := []string{
		"lowscore(bound=10)+usersample(rate=0.1,enabled=true)+mutilple(a=1,b=2)+single()+testing(c=1,d=2)",
		"lowscore(bound=10)+usersample(rate=0.1,enabled=true)+mutilple(a=1,b=2)+single()",
		"lowscore(bound=10)",
		"usersample(rate=0.1,enabled=true)",
		"mutilple(a=1,b=2)+single()",
	}

	for _, expr := range expressions {
		fmt.Printf("==%v==\n", expr)
		if matches := re.FindAllStringSubmatch(expr, -1); matches != nil {
			for i, match := range matches {
				fmt.Printf("==%v== %dth match: %v\n", expr, i, match)
				var e Expression
				if match[1] != "" {
					e.FunctionName = match[1]
				}
				if match[2] != "" && match[3] != "" {
					e.Parameters = append(e.Parameters, Parameter{
						Name:  match[2],
						Value: match[3],
					})
				}
				if match[4] != "" && match[5] != "" {
					e.Parameters = append(e.Parameters, Parameter{
						Name:  match[4],
						Value: match[5],
					})
				}
				fmt.Printf("Expression: %s\n", e.FunctionName)
				for _, param := range e.Parameters {
					fmt.Printf("  - %s = %s\n", param.Name, param.Value)
				}
				fmt.Println()
			}
		} else {
			fmt.Printf("Expression '%s' does not match the regex\n", expr)
		}
	}
}

type MyError struct {
	code string
	msg  string
}

func (e *MyError) Error() string {
	return e.msg
}

func (e *MyError) Code() string {
	return e.code
}

func (e *MyError) WithArgs(a ...interface{}) *MyError {
	e2 := *e
	e2.msg = fmt.Sprintf(e.msg, a...)
	return &e2
}

var ErrWrongParameterType = &MyError{"WrongParameterType", "arg %d[%s] expect a(n) %v type, but got a(n) %v, which is considered as %s"}
