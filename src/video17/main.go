package main

import (
	"fmt"
	"os"
)

// 使用“面向对象”的思维方式编写一个学生信息管理系统
// 1.添加学员信息
// 2，编辑学员信息
// 3。显示学员信息

func showMenu() {
	fmt.Println("欢迎使用学员管理系统")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.显示所有学员信息")
	fmt.Println("4.退出")
}

// 获取用户输入的学员信息
func getInput() *student {
	var (
		id          int
		name, class string
		age, score  uint8
	)
	fmt.Println("请按要求输入：")
	fmt.Print("请输入学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入年龄：")
	fmt.Scanf("%d\n", &age)
	fmt.Print("请输入班级：")
	fmt.Scanf("%s\n", &class)
	fmt.Print("请输入分数：")
	fmt.Scanf("%d\n", &score)
	stu := newStudent(id, name, class, age, score)
	return stu
}

func main() {
	sm := newStudentMgr()
	for {
		// 1。打印系统菜单
		showMenu()
		// 2，等待用户选择操作
		var input int
		fmt.Print("请输入操作：")
		fmt.Scanf("%d\n", &input)
		// 3。执行用户选择的操作
		switch input {
		case 1:
			// 添加学员
			tmpStu := getInput()
			sm.addStudent(tmpStu)
		case 2:
			// 编辑学员
			tmpStu := getInput()
			sm.modifyStudent(tmpStu)
		case 3:
			// 显示学员
			sm.showStudents()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("输入错误，请重新输入。")
		}
	}
}
