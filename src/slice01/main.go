package main

import "fmt"

// 切片slice
func main() {
	//var a []string
	//var b []int
	//var c = []bool{true, false} // 定义同时初始化
	//fmt.Println(a, b, c)

	//// 基于数组定义切片
	//d := [5]int{55, 56, 57, 58, 59}
	//e := d[1:4]
	//fmt.Println(e)        // 【56 57 58】
	//fmt.Printf("%T\n", e) // slice类型

	//// 切片切片
	//f := e[0:2]
	//fmt.Println(f)
	//fmt.Printf("%T\n", f)

	//// make函数构造切片
	//g := make([]int, 5, 10)
	//fmt.Printf("%T\n", g)

	//// 通过len和cap获取切片长度和容量
	//fmt.Println(len(g))
	//fmt.Println(cap(g))
	//// nil
	//var h []int// 声明int切片
	//var i = []int{}// 声明并且初始化
	//var j = make([]int, 0, 0)
	//if a == nil {
	//	fmt.Println("h is a nil")
	//}
	//fmt.Println(h, len(h), cap(h))
	//if i == nil {
	//	fmt.Println("i is a nil")
	//} else {
	//	fmt.Println("i is not a nil")
	//}
	//fmt.Println(i, len(i), cap(i))
	//if j == nil {
	//	fmt.Println("j is a nil")
	//} else {
	//	fmt.Println("j is not a nil")
	//}
	//fmt.Println(j, len(j), cap(j))

	//// 切片的赋值拷贝，传递的是指针
	//k := make([]int, 3)
	//l := k // l,k共享同一个底层数组
	//k[0] = 100 // 同时影响到l[0]
	//fmt.Println(k, l)

	//// 切片的遍历
	//m := []int{1,2,3,4,5}
	////根据索引遍历
	//for i:=0; i<len(m); i++ {
	//	fmt.Println(i, m[i])
	//}
	//fmt.Println()
	//// for range遍历
	//for index, value := range m {
	//	fmt.Println(index, value)
	//}

	//// append方法向切片添加元素（切片扩容）
	//var n []int // n没有申请内存
	//n = append(n, 10)
	//fmt.Println(n)
	//for i:=0; i<10; i++ {
	//	n = append(n, i)
	//	fmt.Printf("%v len:%d cap:%d ptr:%p\n", n, len(n), cap(n), n)
	//}
	//n = append(n, 1,2,3,4,5,6,7,8,9)
	//fmt.Println(n)
	//o := []int{11,12,13,14}
	//n = append(n, o...)
	//fmt.Println(n)

	//// copy函数复制切片
	//p := []int{1,2,3,4,5}
	//q := make([]int, 5, 5)
	//r := q
	//copy(q,p)
	//q[0] = 100
	//fmt.Println(p,q,r)

	// 删除切片元素
	s := []string{"beijing", "shanghai", "guangzhou", "shenzhen"}
	s = append(s[0:2], s[3:]...)
	fmt.Println(s)
}
