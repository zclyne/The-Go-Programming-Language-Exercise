package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(string(reverseUTF8([]byte("Hello, World"))))
	fmt.Println(string(reverseUTF8([]byte("Hello, 世界"))))
}

func reverseUTF8(bytes []byte) []byte {
	if len(bytes) <= 1 { // no need to reverse
		return bytes
	}
	_, d := utf8.DecodeRune(bytes) // get the length of the 1st utf-8 rune in bytes
	// reverse the remaining parts of bytes, and then append the 1st rune to the tail of bytes
	return append(reverseUTF8(bytes[d:]), bytes[:d]...)
}