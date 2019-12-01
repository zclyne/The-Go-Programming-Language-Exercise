// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	var typesCount = []int{0, 0, 0} // 0 for letter, 1 for digit, 2 for others
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF || r == '\n' {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		fmt.Println("Get rune: ", r)
		if unicode.IsLetter(r) {
			typesCount[0]++
		} else if unicode.IsDigit(r) {
			typesCount[1]++
		} else {
			typesCount[2]++
		}
	}
	fmt.Printf("%d letters, %d digits, and %d others\n", typesCount[0], typesCount[1], typesCount[2])
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}