package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段的可见性和JSON序列化

// 定义一个标识符首字母大写，那么对这个包外可见

type Student struct {
	ID   int
	Name string
}

type class struct {
	Title    string    `json:"title"`
	Students []Student `json:"students_list" xml:"ss"` // 指定了在其他解析中使用的Tag
}

// 构造函数
func newStudent(id int, name string) Student {
	return Student{
		ID:   id,
		Name: name,
	}
}
func main() {
	// 创建一个班级变量
	c1 := class{
		Title:    "火箭101",
		Students: make([]Student, 0, 20),
	}
	// 向班级中添加学生
	for i := 0; i < 10; i++ {
		// 造10个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	//fmt.Printf("%#v\n", c1)
	// json的序列化：把Go语言中的数据，转化成json格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal fail! err:", err)
		return
	}
	fmt.Printf("%s\n", data)
	fmt.Printf("%T\n", data)

	//json反序列化：逆向转化
	jsonStr := `{"title":"火箭101","students_list":[{"ID":0,"Name":"小王子"},{"ID":1,"Name":"娜扎"}]}`
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("fail! err:", err)
		return
	}
	fmt.Printf("%#v\n", c2)
	fmt.Printf("%T\n", c2)
}
