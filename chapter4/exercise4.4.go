package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(rotate(a, 2))
	fmt.Println(rotate(a, 0))
	fmt.Println(rotate(a, 8))
}

// rotate() rotates the slice a left by n elements, and returns the result
func rotate(a []int, n int) []int {
	if len(a) == 0 {
		return a
	}
	n %= len(a) // if n is greater than len(a), the effect is equivalent to rotate by n = n % len(a)
	for n > 0 {
		tmp := a[0]
		copy(a[:len(a) - 1], a[1:])
		a[len(a) - 1] = tmp
		n--
	}
	return a
}