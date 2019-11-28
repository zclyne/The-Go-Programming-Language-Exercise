package main

import "fmt"

func main() {
	fmt.Println(PopCount3(1346347))
}

// PopCount3 returns the population count (number of set bits) of x.
func PopCount3(x uint64) int {
	result := 0
	for ; x > 0; x = x & (x - 1) {
		result++
	}
	return result
}