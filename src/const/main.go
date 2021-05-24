package main

import "fmt"

//常量
//const pi = 3.14159
//const e = 2.7

const (
	pi = 3.14
//	e = 2.7
)

const (
	n1 = 10
	n2
	n3
)

const (
	m1 = iota //0
	m2        //1
	m3        //2
	_         //忽略
	m5        //4
)

const (
	x1 = iota //0
	x2        //1
	x3 = 100  //100
	x4        //100
	x5 = iota //4
	x6        //5
)

const (
	_ = iota
	KB = 1<<(10 * iota)//1<<10
	MB                 //1<<20
	GB                 //1<<30
)

const (
	a, b = iota + 1, iota + 2//iota=0, 1, 2
	c, d                     //iota=1, 2, 3
	e, f                     //iota=2, 3, 4
)

func main() {
	fmt.Println(n1, n2, n3)
	fmt.Println(m1, m2, m3, m5)
	fmt.Println(x1,x2,x3,x4,x5,x6)
	fmt.Println(KB, MB, GB)
	fmt.Println(a, b, c, d, e, f)
}
