package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/traefik/yaegi/interp"
)

func main() {
	var err error
	i := interp.New(interp.Options{})
	type Super struct {
	}
	imports := map[string]reflect.Value{
		"Super": reflect.ValueOf((*Super)(nil)),
	}
	_ = i.Use(interp.Exports{
		"fake.com/engine/proto/.": imports,
	})

	v, err := i.Eval(`
		func Playing1() func(any) (any) {
			return func(any) (any) {
				return 1
			}
		}
`)
	if err != nil {
		log.Fatal("compile playing1 err: %v", err)
	}
	if v, err = i.Eval("Playing1"); err != nil {
		log.Fatalf("get the closure generated by Playing1 err: %v", err)
	}
	vv := v.Call(nil)
	x := vv[0].Interface().(func(any) any)
	_ = x(123)
	fmt.Println("finished")
}