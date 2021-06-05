package main

import "fmt"

type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams1(dreams []string) {
	p.dreams = dreams
}

func (p *Person) SetDreams2(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func main() {
	p1 := Person{name: "小王子", age: 18}
	p2 := Person{name: "娜扎", age: 20}
	fmt.Println(p1, p2)
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams1(data)
	p2.SetDreams2(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // 切片共享同一底层数组，修改data[]会影响p1.dreams
	fmt.Println(p2.dreams)
}
