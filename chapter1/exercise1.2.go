package main

import (
	"fmt"
	"os"
	"strconv"
)

// 在terminal当前文件夹中输入go run exercise1.2.go abc def
// 输出结果为：
// 1. abc  2. def

func main() {
	s := ""
	for i, arg := range os.Args[1:] {
		s += strconv.Itoa(i + 1) + ". " + arg + "\t"
	}
	fmt.Println(s)
}
