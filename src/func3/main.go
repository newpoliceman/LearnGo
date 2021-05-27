package main

import "fmt"

//定义一个返回值为函数类型的函数
func a(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}
func main() {
	// 闭包 = 函数 + 外层变量的饮用
	r := a("沙河娜扎") // r此时是一个闭包
	r()
}
