package main

import "fmt"

// 结构体指针

type person struct {
	name, city string
	age        int8
}

func main() {
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	p2.name = "沙河小王子" // p2是指向结构体的指针，在Go语言中，允许使用p2.name，相当于 (*p2).name
	p2.age = 10
	fmt.Printf("%#v\n", p2)

	// 取结构体的地址进行实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	p3.name = "沙河娜扎"
	p3.city = "Beijing"
	p3.age = 18
	fmt.Printf("%#v\n", p3)
}
