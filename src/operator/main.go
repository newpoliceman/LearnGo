package main

import "fmt"

//Go语言中的运算符
func main() {
	//1.算数运算符
	a := 5
	b := 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //++不是运算符，是语句
	fmt.Println(a)
	b-- //--不是运算符，是语句
	fmt.Println(b)
	//2.关系运算符
	fmt.Println(10 > 2)
	fmt.Println(10 != 2)
	fmt.Println(4 <= 5)
	//3.逻辑运算符
	fmt.Println(10 > 5 && 1 == 1)
	fmt.Println(10 > 5 || 1 == 2)
	fmt.Println(!(10>5))
	//4.位运算符
	c := 1//001
	d := 5//101
	fmt.Printf("%b\n", c & d)//001
	fmt.Printf("%b\n", c | d)//101
	fmt.Printf("%b\n", c ^ d)//100
	fmt.Println(1 << 2)//100 4
	fmt.Println(4 >> 2)//001 1
	fmt.Println(1 << 10)//1024
	//5.赋值运算符
	e := 5
	e += 5 // e = e + 5
	fmt.Println(e)
}
