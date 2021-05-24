package main

import "fmt"

//switch判断
func main() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小指")
	default:
		fmt.Println("无效输入")
	}

	//2.case判断多个值
	num := 5
	switch num {
	case 1, 3, 5, 7, 9:
		fmt.Println("odd")
	case 2, 4, 6, 8, 10:
		fmt.Println("even")
	}
	//3.case中使用条件表达式
	age := 18
	switch {
	case age > 18:
		fmt.Println("成年")
	case age < 18:
		fmt.Println("未成年")
	default:
		fmt.Println("???")
	}

}
