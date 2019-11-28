// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 实现方法：把counts的值类型从int改成[]string，存行时，不把出现次数作为值
// 而是把这一行所属的的文件的文件名数组作为值，某一行重复出现时，这个数组的长度就会超过1
// 最后遍历map时，若发现某一行对应的值的长度超过1，就表明这一行重复出现

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, filenames := range counts {
		if len(filenames) > 1 { // has duplicate
			fmt.Printf("%d\t%s\n", len(filenames), line)
			for _, filename := range filenames {
				fmt.Printf("%s\t", filename)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// 在值的slice尾部添加上当前文件名
		counts[input.Text()] = append(counts[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}