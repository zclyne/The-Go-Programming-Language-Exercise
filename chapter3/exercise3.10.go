package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123456"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("123"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n, mod := len(s), len(s) % 3
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
		if (i + 1) % 3 == mod && i != len(s) - 1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}