package main

import "fmt"

//for循环
func main() {
	//1。基本写法
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//2。省略初始语句
	var i = 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	//3。省略初始语句和结束语句，类似while
	var n = 10
	for n > 0 {
		fmt.Println(n)
		n--
	}
	//4。无限循环
	//for {
	//	fmt.Println("hello沙河")
	//}
	//5.break 跳出for循环
	//for m := 0; m < 5; m++ {
	//	fmt.Println(m)
	//	if m == 3 {
	//		break
	//	}
	//}
	//6。continue 继续下一次循环
	for m := 0; m < 5; m++ {
		if m == 3 {
			continue//跳过本次for，继续下一次
		}
		fmt.Println(m)
	}
}
