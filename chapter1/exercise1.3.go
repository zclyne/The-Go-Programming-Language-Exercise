package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 由于命令行参数数量少，所以输出都是0ms，没有差异

func main() {
	// basic version
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%dms elapsed for basic version\n", time.Since(start).Milliseconds())

	// strings.Join version
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%dms elapsed for strings.Join version\n", time.Since(start).Milliseconds())
}
