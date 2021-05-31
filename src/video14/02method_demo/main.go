package main

import "fmt"

// 为任意类型定义方法
// 造一个自己的类型
type myInt int

func (m myInt) sayHi() {
	fmt.Printf("Hi!")
}
func main() {
	var m1 = new(myInt)
	*m1 = 100
	m1.sayHi()
}
