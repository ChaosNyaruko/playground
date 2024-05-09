package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/traefik/yaegi/interp"
)

type HookContainer[T any] struct {
	inner T
	args  []any
}

// func (h HookContainer[T]) Compile() int {
// 	return h.inner(h.args...)
// }

type StreamInfo struct {
	Concurrent int
}
type ClientInfo struct {
	Platform int
}

func ParseComponentT[T any, P any](content string, name string, args ...any) (*HookContainer[T], P, error) {
	var err error
	var res T
	var out P
	r := reflect.TypeOf(res)
	log.Printf("%v, %v, name:%v", r, r.Kind(), r.Name())
	if r.Kind() != reflect.Func {
		return nil, out, fmt.Errorf("the res must be a func type, but got %v", r.Kind())
	}
	// res = reflect.New(reflect.TypeOf(res)).Interface().(T)
	i := interp.New(interp.Options{})
	if err = i.Use(interp.Exports{
		"fake.com/engine/prototype/.": {
			"StreamInfo": reflect.ValueOf((*StreamInfo)(nil)),
			"ClientInfo": reflect.ValueOf((*ClientInfo)(nil)),
		},
	}); err != nil {
		return &HookContainer[T]{inner: res, args: args}, out, err
	}

	if _, err = i.Eval(content); err != nil {
		log.Printf("eval err: %v, err")
		return &HookContainer[T]{inner: res, args: args}, out, err
	}
	updateFunc(&res, i, name)

	// TODO: check correctness
	in := []reflect.Value{}
	for i := 0; i < len(args); i++ {
		in = append(in, reflect.ValueOf(args[i]))
	}
	f := reflect.ValueOf(res).Call(in)
	out = f[0].Interface().(P)
	return &HookContainer[T]{inner: res, args: args}, out, nil
}

type ComponentForHook1 = func(s *StreamInfo, c *ClientInfo) (int, int, int, int)

var c1 = `
	import (
		. "fake.com/engine/prototype"
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
		. "fake.com/engine/prototype"
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

func updateFunc[T any](f *T, i *interp.Interpreter, name string) {
	v, err := i.Eval(name)
	log.Printf("interp.Eval err: %v", err)
	if err == nil {
		*f = v.Interface().(T)
	}
}

func main() {
	hc1, c1, err := ParseComponentT[WrappedComponent1ForHook1, ComponentForHook1](c1, "WrappedC1", 50)
	if err != nil {
		panic(err)
	}
	_ = hc1.inner(hc1.args[0].(int))
	r := reflect.TypeOf(c1)
	log.Printf("c1: %v, %v, name:%v", r, r.Kind(), r.Name())

	hc2, c2, err := ParseComponentT[WrappedComponent2ForHook1, ComponentForHook1](c2, "WrappedC2")
	if err != nil {
		panic(err)
	}
	_ = hc2.inner()
	r = reflect.TypeOf(c2)
	log.Printf("c2: %v, %v, name:%v", r, r.Kind(), r.Name())
}

type WrappedComponent1ForHook1 = func(int) ComponentForHook1

// func (w WrappedComponent1ForHook1) Call(i int) ComponentForHook1 {
// 	return w(i)
// }

type WrappedComponent2ForHook1 = func() ComponentForHook1

// func (w WrappedComponent2ForHook1) Call() ComponentForHook1 {
// 	return w()
// }
