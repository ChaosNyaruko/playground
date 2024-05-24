package engine_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ChaosNyaruko/playground/simplego/engine"
)

// func TestCheckContent(t *testing.T) {
// 	f := `
// 	func WrappedC1(bound int) func(*StreamInfo, *ClientInfo) (int, int, int, int) {
// 		x := bound
// 		return func(s *StreamInfo, c *ClientInfo) (id, reset, code, quit int) {
// 			if s.Concurrent > x {
// 				return 2003, 1, 233, 1
// 			}
// 			if c.Platform == 4 {
// 				return 2002, 0, 234, 1
// 			}
// 			return 2002, 0, 235, 0
// 		}
// 	}`
// 	ast, err := engine.CheckContent(f)
// 	t.Logf("ast: %#v, err: %v", ast, err)
// 	assert.Nil(t, err)
// }

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

func TestParse(t *testing.T) {
	f, err := engine.OldParseComponent(c1, "WrappedC1", 50)
	assert.Nil(t, err)
	assert.NotNil(t, f)

	c1 := f // WrappedComponent(50)(*StreamInfo, *ClientInfo) (int, int, int, int)

	x, y, z, quit := c1(&engine.StreamInfo{Concurrent: 60}, &engine.ClientInfo{Platform: 1})
	assert.Equal(t, 2003, x)
	assert.Equal(t, 1, y)
	assert.Equal(t, 233, z)
	assert.Equal(t, quit, 1)

	x, y, z, quit = c1(&engine.StreamInfo{Concurrent: 20}, &engine.ClientInfo{Platform: 4})
	assert.Equal(t, 2002, x)
	assert.Equal(t, 0, y)
	assert.Equal(t, 234, z)
	assert.Equal(t, quit, 1)

	x, y, z, quit = c1(&engine.StreamInfo{Concurrent: 20}, &engine.ClientInfo{Platform: 1})
	assert.Equal(t, 2002, x)
	assert.Equal(t, 0, y)
	assert.Equal(t, 235, z)
	assert.Equal(t, quit, 0)
}

func TestCompose(t *testing.T) {
	// Rough prototype, no proper module abstractions/interfaces, just for demonstration.

	// ------- engine.ParsePipeline("1(bound=50)+2") 语义
	//
	// douyin-ios: 1(bound=50)+2
	// douyin-android: 2()+1(bound=50)
	// xigua-ios: 1(bound=100)+2 not showed here
	f, err := engine.OldParseComponent(c1, "WrappedC1", 50)
	assert.Nil(t, err)
	g, err := engine.OldParseComponent(c2, "WrappedC2") // WrappedC2 is a component that without any arguments.
	assert.Nil(t, err)

	// ------ TODO: who should do this registration?
	m := map[string]engine.PipelineForHook1{} // TODO: can be an interface, with { Get(a, b, c) Executor}
	m["douyin-ios"] = engine.Compose([]engine.ComponentForHook1{f, g})
	m["douyin-android"] = engine.Compose([]engine.ComponentForHook1{g, f})

	x := &engine.StreamInfo{Concurrent: 137}
	y := &engine.ClientInfo{Platform: 1}

	// ---service hook calling
	// f hits the first condition, and quits, g has no effect.
	id, reset, code := m["douyin-ios"].Execute(x, y)
	assert.Equal(t, 2003, id)
	assert.Equal(t, 1, reset)
	assert.Equal(t, 233, code)

	// g hits the first condition, and quits, f has no effect.
	id, reset, code = m["douyin-android"].Execute(x, y)
	assert.Equal(t, 2003, id)
	assert.Equal(t, 1, reset)
	assert.Equal(t, 300, code)

	// f misses all its conditions, but returns quit=0, so continue g
	x.Concurrent = 0
	id, reset, code = m["douyin-ios"].Execute(x, y)
	assert.Equal(t, 2002, id)
	assert.Equal(t, 0, reset)
	assert.Equal(t, 235, code)
}

func TestRuntimeGroupID(t *testing.T) {
	err := engine.Init()
	assert.Nil(t, err)

	a := &engine.StreamInfo{Concurrent: 137}
	b := &engine.ClientInfo{Platform: 1}
	x := &engine.Mixed{
		a,
		b,
	}

	// douyin ios
	// f hits the first condition, and quits, g has no effect.
	id, reset, code := engine.GetRuntimeGroupID1(x, 1128, 4, 0, "")
	assert.Equal(t, 2003, id)
	assert.Equal(t, 1, reset)
	assert.Equal(t, 233, code)

	// f misses all its conditions, but returns quit=0, so continue g
	x.Concurrent = 0
	id, reset, code = engine.GetRuntimeGroupID1(x, 1128, 4, 0, "")
	assert.Equal(t, 2002, id)
	assert.Equal(t, 0, reset)
	assert.Equal(t, 235, code)

	// douyin-android
	// g hits the first condition, and quits, f has no effect.
	x.Concurrent = 137
	// id, reset, code = engine.GetRuntimeGroupID(x, y)
	id, reset, code = engine.GetRuntimeGroupID1(x, 1128, 1, 0, "")
	assert.Equal(t, 2003, id)
	assert.Equal(t, 1, reset)
	assert.Equal(t, 300, code)
}

func TestParseP(t *testing.T) {
	// Test the expressions
	type testcase struct {
		input  string
		expect error
	}
	testcases := []testcase{
		// {"WrappedC1(bound=50)+WrappedC2()", nil},
		{"WrappedC1(bound=\"bug\")+WrappedC2()", engine.ErrWrongParameterType.WithArgs(0, "bound", "int", `"bug"`, `string`)},
		{"WrappedC1(bound=false)+WrappedC2()", engine.ErrWrongParameterType.WithArgs(0, "bound", "int", `false`, "bool")},
		{"WrappedC1(bound=0.1)+WrappedC2()", engine.ErrWrongParameterType.WithArgs(0, "bound", "int", `0.1`, "float64")},
		{"WrappedAny(bound=0.1)+WrappedC2()", nil},
		// "single()",
		// "lowscore(bound=10)",
		// "usersample(rate=0.1,enabled=true)",
		// "mutilple(a=1,b=2)+single()",
		// "lowscore(bound=10)+usersample(rate=0.1,enabled=true)+mutilple(a=1,b=2)+single()",
		// "lowscore(bound=10)+usersample(rate=0.1,enabled=true)+mutilple(a=1,b=2)+single()+testing(a=1,b=2,c=3,d=4)",
	}
	for _, e := range testcases {
		_, err := engine.ParseP(e.input, "")
		assert.Equal(t, e.expect, err)
	}
}

func TestX(t *testing.T) {
	engine.X()
}
