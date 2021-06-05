package main

import "fmt"

// 结构体继承
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动。\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal // 匿名嵌套，而且嵌套了一个指针
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪。\n", d.name)
}
func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "Lele",
		},
	}
	fmt.Printf("%v\n", d1)
	d1.wang() //
	d1.move() // Dog继承了嵌套结构体Animal的move()方法
}
