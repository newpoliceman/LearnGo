package main

import "fmt"

// 定义结构体
type person struct {
	name, city string
	age        int8
}

func main() {
	// 结构体实例化
	var p person
	p.name = "沙河娜扎"
	p.city = "Beijing"
	p.age = 18
	fmt.Printf("p=%#v\n", p)

	//匿名结构体
	var user struct {
		name     string
		age      int8
		marriage bool
	}
	user.marriage = false
	user.name = "沙河娜扎"
	user.age = 18

}
