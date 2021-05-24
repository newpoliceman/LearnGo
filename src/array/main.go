package main

import "fmt"

// 数组相关
func main() {
	//var a [3]int
	//var b [4]int
	//fmt.Println(a)
	//fmt.Println(b)
	//// 数组的初始化
	//// 1。使用初始值列表的方式初始化
	//var cityArray = [4]string{"beijing", "shanghai", "guangzhou", "shenzhen"}
	//fmt.Println(cityArray)
	//fmt.Println(cityArray[3])
	//// 2.编译器推导数组容量
	//var boolArray = [...]bool{true, true, false, true}
	//fmt.Println(len(boolArray)) // 4
	//// 3.使用索引值的方式初始化数组
	//var langArray = [...]string{1: "Golang", 2: "Python", 7: "Java"}
	//fmt.Println(len(langArray))   // 8
	//fmt.Printf("%T\n", langArray) // [8]string
	//
	//// 数组的遍历
	//// 1。使用for循环
	//for i := 0; i < len(cityArray); i++ {
	//	fmt.Println(cityArray[i])
	//}
	//// 2.使用for range遍历
	//for index, value := range cityArray {
	//	fmt.Println(index, value)
	//}
	// 二维数组
	cityArray := [...][2]string{
		{"beijing", "xi'an"},
		{"shanghai", "hangzhou"},
		{"chengdu", "chongqing"},
		{"guangzhou", "shenzhen"},
	}
	//fmt.Println(cityArray)
	//fmt.Println(cityArray[2][1])
	// 二维数组的遍历
	for _, v1 := range cityArray {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	// 数组是值类型，赋值给另一个变量时复制的是值的拷贝，修改被赋值的变量原来的变量不变
	x := [3][2]int{
		{1, 2},
		{3,4},
		{5,6},
	}
	fmt.Println(x) // [[1 2] [3 4] [5 6]]
	f1(x)
	fmt.Println(x) // [[1 2] [3 4] [5 6]]
	y := x
	y[0][0] = 1000
	fmt.Println(x) // [[1 2] [3 4] [5 6]]
	fmt.Println(y) // [[1000 2] [3 4] [5 6]]

}

func f1(a [3][2]int) {
	a[0][0] = 100
}
