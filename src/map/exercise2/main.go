package main

import "fmt"

func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\nptr:%p\n", s, s)
	fmt.Printf("%+v\nptr:%p\n", m["q1mi"], m["q1mi"])
}
