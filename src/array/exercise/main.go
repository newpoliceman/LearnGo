package main

import "fmt"

func main() {
	// 练习1
	var array = [...]int{1, 3, 5, 7, 8}
	fmt.Println(array)
	sum := 0
	for _, v := range array {
		sum += v
	}
	fmt.Println("Sum is", sum)

	//练习2
	for index, value := range array {
		for j := len(array)-1; j > index; j-- {
			if (value + array[j]) == 8 {
				fmt.Printf("(%d, %d)\n", index, j)
			}
		}
	}

}
