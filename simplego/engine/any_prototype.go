package engine

var cany_ = `
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
var c1_ = `
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

var c2_ = `
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

type Mixed struct {
	*StreamInfo
	*ClientInfo
}
