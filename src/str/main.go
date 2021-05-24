package main

import (
	"fmt"
	"strings"
)

//字符串操作
func main() {
	//求串长
	s := "hello"
	fmt.Println(len(s))
	s2 := "hello沙河"
	fmt.Println(len(s2))

	//拼接字符串
	fmt.Println(s + s2)
	s3 := fmt.Sprintf("%s - %s", s, s2)
	fmt.Println(s3)

	//切割字符串
	s4 := "how do you do"
	fmt.Println(strings.Split(s4, " "))
	fmt.Printf("%T\n", strings.Split(s4, " "))
	//判断包含
	fmt.Println(strings.Contains(s4, "do"))
	//判断前缀，后缀
	fmt.Println(strings.HasPrefix(s4, "how"))
	fmt.Println(strings.HasSuffix(s4, "how"))
	//判断字串位置
	fmt.Println(strings.Index(s4, "do"))
	//判断字串最后出现位置
	fmt.Println(strings.LastIndex(s4, "do"))
	//字符串切片连接
	s5 := []string{"how", "do", "you", "do"}
	fmt.Println(s5)
	fmt.Println(strings.Join(s5, "+"))


}
