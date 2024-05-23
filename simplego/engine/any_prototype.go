package engine

import "fmt"

var cany_ = `
	import (
		. "fake.com/engine/proto"
	)

	func WrappedAny(bound interface{}) func(any) (int, int, int, int) {
		x := bound.(int)
		return func(s any) (id, reset, code, quit int) {
			if s.(*Mixed).Concurrent > x {
				return 2003, 1, 233, 1
			}
			if s.(*Mixed).Platform == 4 {
				return 2002, 0, 234, 1
			}
			return 2002, 0, 235, 0
		}
	}
	`
var c1_ = `
	import (
		. "fake.com/engine/proto"
	)

	func WrappedC1(bound int) func(any) (int, int, int, int) {
		x := bound
		return func(s any) (id, reset, code, quit int) {
			if s.(*Mixed).Concurrent > x {
				return 2003, 1, 233, 1
			}
			if s.(*Mixed).Platform == 4 {
				return 2002, 0, 234, 1
			}
			return 2002, 0, 235, 0
		}
	}
`

var c2_ = `
	import (
		. "fake.com/engine/proto"
	)

	func WrappedC2() func(any) (int, int, int, int) {
		return func(s any) (id, reset, code, quit int) {
			if s.(*Mixed).Concurrent == 137 {
				return 2003, 1, 300, 1
			}
			return 2002, 0, 235, 0
		}
	}
`

type Mixed struct {
	*StreamInfo
	*ClientInfo
}

var anyStub = func(any) (int, int, int, int) {
	return 0, 0, 0, 0
}

type ComponentForAny = func(any) (int, int, int, int)

func ComposeAny(components []ComponentForAny) PipelineForAny {
	return PipelineForAny{
		Components: components,
	}
}

type PipelineForAny struct {
	Components []ComponentForAny
}

func (p PipelineForAny) Execute(s any) (id, reset, code int) {
	quit := 0
	for _, checker := range p.Components {
		id, reset, code, quit = checker(s)
		if quit != 0 {
			return
		}
	}
	return
}

type RouterAny struct {
	normal   map[int]map[int]PipelineForAny
	override map[int]PipelineForAny
}

func (r *RouterAny) Init() error {
	r.normal = make(map[int]map[int]PipelineForAny)
	r.override = make(map[int]PipelineForAny)
	return nil
}

func (r *RouterAny) Register(t string, sid, appID, clientType int, p PipelineForAny) error {
	switch t {
	case "override":
		r.override[sid] = p
	case "normal":
		if _, ok := r.normal[appID]; !ok {
			r.normal[appID] = make(map[int]PipelineForAny)
		}
		r.normal[appID][clientType] = p
	default:
		return fmt.Errorf("unexpected type for register: %v", t)
	}
	return nil
}

func (r *RouterAny) Get(appid int, clienttype int, sid int, extra string) PipelineForAny {
	// TODO: just a demo, no error handling yet
	if p, ok := r.override[sid]; ok {
		return p
	}
	return r.normal[appid][clienttype]
}

type HookAny struct {
	router RouterAny
}

var A *HookAny = &HookAny{}

func (h *HookAny) Init() {
	_ = h.router.Init()
}

func GetRuntimeGroupIDAny(s any, appid int, clienttype int, sid int, extra string) (int, int, int) {
	return A.router.Get(appid, clienttype, sid, extra).Execute(s)
}

var anyWrapperInt = func(int) ComponentForAny {
	return func(any) (int, int, int, int) {
		return 0, 0, 0, 0
	}
}

var anyWrapperEmpty = func() ComponentForAny {
	return func(any) (int, int, int, int) {
		return 0, 0, 0, 0
	}
}

var anyWrapperAny = func(any) ComponentForAny {
	return func(any) (int, int, int, int) {
		return 0, 0, 0, 0
	}
}
