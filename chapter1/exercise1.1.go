package main

import (
	"fmt"
	"os"
	"strings"
)

// 在terminal当前文件夹中输入go run exercise1.1.go abc def
// 输出结果为：
// command name: /var/folders/vs/lgj2k6y94l31tvnk8l_k8wzr0000gn/T/go-build700768339/b001/exe/exercise1.1
// abc def

func main() {
	fmt.Printf("command name: %s\n", os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
