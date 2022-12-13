package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	// This initializes gpython for runtime execution and is essential.
	// It defines forward-declared symbols and registers native built-in modules, such as sys and time.
	_ "github.com/go-python/gpython/stdlib"

	// Commonly consumed gpython
	"github.com/go-python/gpython/py"
	"github.com/go-python/gpython/repl"
	"github.com/go-python/gpython/repl/cli"
)

func main() {
	flag.Parse()
	if err := envPython3("testdata/date_time.py"); err != nil {
		log.Println("envPython3 err", err)
		os.Exit(2)
	}
	// runWithFile(flag.Arg(0))
	// if err := runWithSrc(src); err != nil {
	// 	log.Println(err)
	// }
}

var src = `
a = "abc"; print(a)
print("hh")
print("hh")
print("hh")
print("hh")
`

func envPython3(file string) error {
	python3 := exec.Command("python3", file)
	stdout, err := python3.StdoutPipe()
	if err != nil {
		return err
	}

	r := bufio.NewReader(stdout)
	if err := python3.Start(); err != nil {
		return err
	}
	if s, err := r.ReadString('\n'); err != nil {
		return err
	} else {
		fmt.Printf("%q", s)
	}
	if err := python3.Wait(); err != nil {
		return err
	}

	return nil
}

func runWithSrc(src string) error {
	// See type Context interface and related docs
	ctx := py.NewContext(py.DefaultContextOpts())
	defer ctx.Close()

	var err error
	_, err = py.RunSrc(ctx, src, "a simple source", nil)
	return err
}

func runWithFile(pyFile string) error {

	// See type Context interface and related docs
	opts := py.DefaultContextOpts()
	opts.SysPaths = append(opts.SysPaths, "/opt/homebrew/Cellar/python@3.10/3.10.8/Frameworks/Python.framework/Versions/3.10/lib/python3.10")
	fmt.Println(opts)
	ctx := py.NewContext(opts)

	// This drives modules being able to perform cleanup and release resources
	defer ctx.Close()

	var err error
	if len(pyFile) == 0 {
		replCtx := repl.New(ctx)

		fmt.Print("\n=======  Entering REPL mode, press Ctrl+D to exit  =======\n")

		_, err = py.RunFile(ctx, "lib/REPL-startup.py", py.CompileOpts{}, replCtx.Module)
		if err == nil {
			cli.RunREPL(replCtx)
		}

	} else {
		_, err = py.RunFile(ctx, pyFile, py.CompileOpts{}, nil)
	}

	if err != nil {
		py.TracebackDump(err)
	}

	return err
}
