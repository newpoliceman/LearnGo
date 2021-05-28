package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

/*
答案：
func f1() int {} 函数返回值没有指明变量名字，所以尽管 return x 都是实际上返回的是一个临时变量，并不是 x ，return x 只是对这个临时变量赋值（即 temp := x; return temp，所以后续 defer 的 x++ 不会影响返回值，所以返回 5
func f2() x int {} 明确返回的是 x，那么就不会有临时变量了，就是返回 x，所以返回 6
func f3() y int {} 明确返回 y，所以 return x 等价于 y = x; return y，所以返回是 5
func f4() x int {} 明确返回的是 x,但是由于传参是值传递，所以 defer 的匿名函数对 x 的修改外部不可见，所以还是返回 5
*/
