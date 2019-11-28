package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func main() {
	fmt.Println(PopCount(1346347))
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	result := 0
	for i := 1; i <= 8; i++ {
		result += int(pc[byte(x)])
		x = x / uint64(i * 256) // 相当于右移8bit
	}
	return result
}