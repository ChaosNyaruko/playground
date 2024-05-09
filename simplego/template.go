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

func ParseComponentT[T any](content string, name string, args ...any) (*HookContainer[T], error) {
	var err error
	var res T
	r := reflect.TypeOf(res)
	log.Printf("%v, %v, name:%v", r, r.Kind(), r.Name())
	if r.Kind() != reflect.Func {
		return nil, fmt.Errorf("the res must be a func type, but got %v", r.Kind())
	}
	// res = reflect.New(reflect.TypeOf(res)).Interface().(T)
	i := interp.New(interp.Options{})
	if err = i.Use(interp.Exports{
		"fake.com/engine/prototype/.": {
			"StreamInfo": reflect.ValueOf((*StreamInfo)(nil)),
			"ClientInfo": reflect.ValueOf((*ClientInfo)(nil)),
		},
	}); err != nil {
		return &HookContainer[T]{inner: res, args: args}, err
	}

	if _, err = i.Eval(content); err != nil {
		log.Printf("eval err: %v, err")
		return &HookContainer[T]{inner: res, args: args}, err
	}
	// numIn := r.NumIn()
	// r.In()
	updateFunc(&res, i, name)
	return &HookContainer[T]{inner: res, args: args}, nil
	// // TODO: better and more robust generics implementation?
	// switch len(args) {
	// case 0:
	// 	var f WrappedComponent2ForHook1
	// 	updateFunc(&f, i, name)
	// 	return f(), err
	// case 1:
	// 	var f WrappedComponent1ForHook1
	// 	updateFunc(&f, i, name)
	// 	return f(args[0].(int)), err
	// default:
	// 	return nil, fmt.Errorf("wrong argument for Parsing components")
	// }
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

func updateFunc[T any](f *T, i *interp.Interpreter, name string) {
	v, err := i.Eval(name)
	log.Printf("interp.Eval err: %v", err)
	if err == nil {
		*f = v.Interface().(T)
	}
}

func main() {
	hc, err := ParseComponentT[WrappedComponent1ForHook1](c1, "WrappedC1", 50)
	if err != nil {
		panic(err)
	}
	c := hc.inner(hc.args[0].(int))
	r := reflect.TypeOf(c)
	log.Printf("%v, %v, name:%v", r, r.Kind(), r.Name())
}

type WrappedComponent1ForHook1 = func(int) ComponentForHook1
type WrappedComponent2ForHook1 = func() ComponentForHook1
