package main

import "fmt"

// 打印乘法表
func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
