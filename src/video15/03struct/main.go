package main

import "fmt"

// 嵌套结构体的字段冲突
type Address struct {
	City       string
	Province   string
	UpdateTime string
}
type Email struct {
	Addr       string
	UpdateTime string
}
type Person struct {
	Name    string
	Gender  string
	Age     int
	Address // 嵌套匿名结构体
	Email
}

func main() {
	p1 := Person{
		Name:   "小王子",
		Gender: "male",
		Age:    18,
		Address: Address{
			Province:   "山东",
			City:       "威海",
			UpdateTime: "2021-02-02",
		},
		Email: Email{
			Addr:       "good@123.com",
			UpdateTime: "2020-01-01",
		},
	}
	// fmt.Printf("%v\n", p1.UpdateTime) // 同名字段冲突,报错
	fmt.Printf("%v\n", p1.Address.UpdateTime)
	fmt.Printf("%v\n", p1.Email.UpdateTime)
}
