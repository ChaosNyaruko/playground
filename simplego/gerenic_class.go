// You can edit this code!
// Click here and start typing.
package main

func main() {
	a := []S{"1", "2"}
	f(a)
}

func f[T I](a []T) {
	//do
	for _, v := range a {
		v.F()
	}
}

type I interface {
	S   //显式注册实现类
	F() //定义接口方法
}
type S string

func (s S) F() {
	println(s)
}
