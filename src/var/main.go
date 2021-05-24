package main

import "fmt"

var x = 100
var y = "小王子"

//Variables
func main() {
	//Standrad Claim
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)
	//批量声明
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)
	//声明同时指定初始值
	var name1 = "小王子"
	var age1 = 18
	fmt.Println(name1, age1)
	var name2, age2 = "沙河娜扎", 28
	fmt.Println(name2, age2)
	//类型推导，编译器自动推导类型
	var name3 = "小王子"
	var age3 = 18
	fmt.Println(name3, age3)
	//简短变量声明
	m := 10
	fmt.Println(m)
	//匿名变量

}