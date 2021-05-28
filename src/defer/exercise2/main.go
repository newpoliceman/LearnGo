package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

/*答案
defer注册要延迟执行的函数时该函数所有的参数都需要确定其值。
第一个defer，为了确定值，运行一次calc("A", x, y)，此时，x=1，y=2，输出A 1 2 3，注册calc("AA", 1, 3)
第二个defer，为了确定值，运行一次calc("B", x, y)，此时，x=10，y=2，输出B 10 2 12，注册calc("BB", 10, 12)
执行y=20后，函数准备执行RET，开始逆序执行注册的defer，输出BB 10 12 22 和 AA 1 3 4
*/
