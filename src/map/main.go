package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// map映射
func main() {
	// 声明而不初始化，初始值是nil
	var a map[string]int
	fmt.Println(a == nil) // true
	// map初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)
	// mao添加键值对
	a["沙河娜扎"] = 100
	a["沙河小王子"] = 200
	fmt.Println(a)
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("type:%T\n", a)
	// 声明同时初始化
	b := map[int]bool{
		1: true,
		2: false,
		4: true,
	}
	fmt.Printf("b:%v\n", b)
	fmt.Printf("type:%T\n", b)
	// 判断键值是否存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["沙河娜扎"] = 100
	scoreMap["沙河小王子"] = 200
	v, ok := scoreMap["张二狗子"]
	if ok {
		fmt.Println("有张二狗子", v)
	} else {
		fmt.Println("查无此人")
	}
	// map的遍历
	// for range
	// map是无序的，与添加顺序无关
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	// 只遍历Key
	for k := range scoreMap {
		fmt.Println(k)
	}
	// 只遍历map的value
	for _, v := range scoreMap {
		fmt.Println(v)
	}
	//delete()函数删除键值对
	delete(scoreMap, "沙河小王子")
	fmt.Println(scoreMap)

	// 按照指定顺序遍历
	var scoreMap2 = make(map[string]int, 100)
	// 添加50个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap2[key] = value
	}
	for k, v := range scoreMap2 {
		fmt.Println(k, v)
	}
	// 按照Key升序遍历map
	// 1.先取出所有的key放入切片
	var keys = make([]string, 0, 100)
	for k := range scoreMap2 {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	// 2.对key进行排序
	sort.Strings(keys)
	// 3.按照排序后的key
	for _, k := range keys {
		fmt.Println(k, scoreMap2[k])
	}

	// 元素类型为map的切片
	var mapSlice = make([]map[string]int, 0, 8) // 仅仅完成了外层切片的初始化
	fmt.Println(mapSlice)
	// 完成内层map的初始化
	mapSlice = append(mapSlice, make(map[string]int, 8))
	mapSlice[0]["沙河小王子"] = 100
	fmt.Println(mapSlice)

	// 值为切片的map
	var sliceMap = make(map[string][]int, 8) // 只完成了外层map的初始化
	fmt.Println(sliceMap)
	value, isok := sliceMap["China"]
	if isok {
		fmt.Println(value)
	} else {
		sliceMap["China"] = make([]int, 8) // 完成了切片的初始化
		sliceMap["China"][0] = 100
		sliceMap["China"][1] = 200
		sliceMap["China"][4] = 300
	}
	//遍历sliceMap
	for k1, v1 := range sliceMap {
		fmt.Println(k1, v1)
	}

}
