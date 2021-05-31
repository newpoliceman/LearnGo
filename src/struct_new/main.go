package main

import "fmt"

// 结构体构造函数：构造一个结构体实例的函数
// 结构体是值类型，每次复制都是完整的值拷贝，内存开销大
type person struct {
	name, city string
	age        int8
}

// 构造函数通常返回一个指针，以节省内存开销
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
func main() {
	p1 := newPerson("小王子", "Beijing", 10)
	fmt.Printf("type:%T value:%#v\n", p1, p1)
}
