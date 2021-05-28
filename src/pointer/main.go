package main

import "fmt"

// 指针
/*
func main() {

	// 变量取地址
	a := 99 // int 整形
	b := &a // *int 整形的指针类型
	fmt.Printf("%v %p\n", a, &a)
	fmt.Printf("%v %p\n", b, &b)
	// 变量b是int类型的指针 *int，存储了变量a的内存地址
	// 指针取值
	c := *b // 根据内存地址取值
	fmt.Println(c) // 99
}

*/
func main() {
	a := 1
	modify(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}

func modify(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}
