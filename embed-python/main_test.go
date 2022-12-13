package main

import (
	"testing"

	"github.com/go-python/gpython/pytest"
)

func TestRunFunction(t *testing.T) {
	pytest.RunScript(t, "./testdata/simple_function.py")
}
