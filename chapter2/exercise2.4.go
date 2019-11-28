package main

import "fmt"

func main() {
	fmt.Println(PopCount2(1346347))
}

// PopCount2 returns the population count (number of set bits) of x.
func PopCount2(x uint64) int {
	result := 0
	for ; x > 0; x = x >> 1 {
		result += int(x) & 1
	}
	return result
}