package main

import "fmt"

type student struct {
	id          int
	name, class string
	age, score  uint8
}

// student类型的构造函数
func newStudent(id int, name, class string, age, score uint8) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
		age:   age,
		score: score,
	}
}

// 学员管理的类型
type studentMgr struct {
	allStudents []*student
}

// 学员管理的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 1。添加学生
func (s *studentMgr) addStudent(tmpStu *student) {
	s.allStudents = append(s.allStudents, tmpStu)
}

// 2.编辑学生
func (s *studentMgr) modifyStudent(tmpStu *student) {
	for i, v := range s.allStudents {
		if tmpStu.id == v.id {
			s.allStudents[i] = tmpStu
			return
		}
	}
	fmt.Printf("系统中没有学号是%d的学生。\n", tmpStu.id)
}

// 3。显示学生
func (s *studentMgr) showStudents() {
	for _, v := range s.allStudents {
		fmt.Printf("学号:%2d, 姓名：%s，年龄：%d，班级：%s，分数：%d。\n", v.id, v.name, v.age, v.class, v.score)
	}
}
