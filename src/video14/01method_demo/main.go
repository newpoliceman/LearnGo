package main

import "fmt"

// 方法的定义实例

// Person 是一个结构体
type Person struct {
	name string
	age  int8
}

// 编写构造函数
func newPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream 为Person类型定义方法
// 值接收者
func (p Person) Dream() {
	fmt.Printf("%d岁的%s的梦想是学好Go语言。\n", p.age, p.name)
}

// SetAge 是一个修改年龄的方法
// 指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2 是一个使用值接收者修改年龄的方法
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := newPerson("沙河娜扎", int8(18))
	p1.Dream()

	// 调用修改年龄的方法
	fmt.Printf("%s原来是%d岁。", p1.name, p1.age)
	p1.SetAge(int8(20))
	fmt.Printf("%s现在是%d岁。", p1.name, p1.age)
	p1.SetAge2(int8(30))
	fmt.Println(p1.age) // 因为方法的接收者是值拷贝，作为值接收者时，不会影响原来的变量，没有修改成30
}
