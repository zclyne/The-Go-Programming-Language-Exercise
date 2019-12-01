package main

import "fmt"

func main() {
	s := [6]int{1, 2, 3, 4, 5, 6}
	reverse(&s)
	fmt.Println(s)
}

// reverse reverses an array of 6 ints in place.
// Although using array pointer enables changing elements in the array
// the size of the array must be specified, therefore it is not practical
func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}