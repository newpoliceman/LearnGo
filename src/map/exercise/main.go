package main

import (
	"fmt"
	"strings"
)

func main() {
	// 统计字符串中每个单词出现的次数
	var s = "how do you do"
	//
	var wordCount = make(map[string]int, 10)
	words := strings.Split(s, " ")
	for _, word := range words {
		_, ok := wordCount[word]
		if ok {
			// map中有这个单词
			wordCount[word] += 1
		} else {
			wordCount[word] = 1
		}
	}

	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}
