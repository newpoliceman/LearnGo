package main

import (
	"fmt"
)

//函数
//定义一个没有参数也没有返回值的函数
func sayHello() {
	fmt.Println("hello, world!")
}

// 定义一个接收string类型的函数
func sayHello2(name string) {
	fmt.Println("hello,", name)
}

// 定义接收多参数函数并且有一个返回值
func intSum(a int, b int) int {
	ret := a + b
	return ret
}

// 定义函数同时声明返回值
func intSum2(a int, b int) (ret int) {
	ret = a + b
	return
}

// 函数接收可变参数，在参数后面加...，表示可变参数，是切片类型，可变参数必须放在最后
// go语言的函数中不存在默认参数
func intSum3(a ...int) int {
	fmt.Printf("%T\n", a)
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

// 定义多返回值的函，多返回值支持类型简写
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
func main() {
	//函数调用
	sayHello()
	name := "沙河娜扎"
	sayHello2(name)
	sayHello2("沙河小王子")
	r := intSum(1, 2)
	fmt.Println(r)
	r2 := intSum3(10, 20, 30)
	fmt.Println(r2)
	x, y := calc(100, 200)
	fmt.Println(x, y)
}
