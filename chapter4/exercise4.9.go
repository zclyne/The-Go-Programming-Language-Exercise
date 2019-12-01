package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCount := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		fmt.Println("Get word: ", word)
		if word == "end" {
			break
		}
		wordCount[word]++
	}
	for word, count := range wordCount {
		fmt.Printf("Word %s appears for %d times\n", word, count)
	}
}
