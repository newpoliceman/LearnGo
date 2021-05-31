package main

import "fmt"

// 结构体的初始化

type person struct {
	name, city string
	age        int8
}

func main() {
	// 1.键值对方式初始化
	p4 := person{
		name: "小王子",
		city: "Beijing",
		age:  10,
	}
	fmt.Printf("%#v\n", p4)

	p5 := &person{
		name: "沙河娜扎",
		city: "Beijing",
		age:  18,
	}
	fmt.Printf("%#v\n", p5)

	// 2.值列表的方式初始化
	p6 := person{
		"小王子",
		"Beijing",
		18,
	}
	fmt.Printf("%#v\n", p6)
}
