package main

import "fmt"

func main() {
	var a *int
	a = new(int) // 用于基础类型内存分配，内存对应值为该类型的0值，返回指向该类型的指针，不常用
	*a = 100
	fmt.Println(*a)

	var b = make(map[string]int, 10) // 用于slice map channel类型的初始化，返回该类型本身
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
