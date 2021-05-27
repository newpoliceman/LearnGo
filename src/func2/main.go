package main

import "fmt"

// 函数进阶之作用域

// 定义全局变量
var num int = 10

//定义函数
func testGlobal() {
	num := 100
	name := "沙河娜扎"
	fmt.Println("变量num", num)
	fmt.Println(name)
}

func main() {
	testGlobal()
	abc := testGlobal
	fmt.Printf("%T\n", abc)
	abc()
}
