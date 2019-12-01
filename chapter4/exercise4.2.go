package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	flag := 0 // 0 for sha256, 1 for sha384, 2 for sha512
	if len(os.Args) != 1 { // no args
		if os.Args[1] == "sha384" || os.Args[1] == "SHA384" {
			flag = 1
		} else if os.Args[1] == "sha512" || os.Args[1] == "SHA512" {
			flag = 2
		} else if os.Args[1] != "sha256" && os.Args[1] != "SHA256" {
			panic("Invalid args! Use sha256/sha384/sha512.")
		}
	}
	// read contents from stdin
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := []byte(scanner.Text())
	if flag == 0 {
		fmt.Printf("%x\n", sha256.Sum256(text))
	} else if flag == 1 {
		fmt.Printf("%x\n", sha512.Sum384(text))
	} else {
		fmt.Printf("%x\n", sha512.Sum512(text))
	}
}
