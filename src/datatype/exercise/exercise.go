package main

import "fmt"

func main() {
	//练习1
	i := 100
	f := 1.0
	b := true
	c := 'a'
	fmt.Printf("i:%T f:%T b:%T c:%T", i, f, b, c)
	//练习2
	s1 := "hello沙河小王子"
	count := 0
	for _,r := range s1{
		if r > 256{
			count++
		}
	}
	fmt.Println()
	fmt.Println("aaaa:", count)
}
