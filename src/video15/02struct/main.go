package main

import "fmt"

// 嵌套结构体
type Person struct {
	Name    string
	Gender  string
	Age     int8
	Address // 嵌套一个结构体类型，匿名结构体
}
type Address struct {
	City     string
	Province string
}

func main() {
	p1 := Person{
		Name:   "小王子",
		Age:    18,
		Gender: "male",
		Address: Address{
			City:     "威海",
			Province: "山东",
		},
	}
	fmt.Println(p1)
	fmt.Printf("%v\n", p1.Address.Province) // 通过嵌套结构体名访问内部字段
	fmt.Println(p1.Province)                // 直接通过内部字段名访问
}
