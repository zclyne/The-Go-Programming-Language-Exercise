package main

import "fmt"

func main() {
	s := []string{"a", "b", "cd", "cd"}
	fmt.Println(removeNeighbourDuplicate(s))
	s = []string{"a", "a", "b", "cd", "cd"}
	fmt.Println(removeNeighbourDuplicate(s))
	s = []string{"a", "a", "a", "b", "cd", "cd"}
	fmt.Println(removeNeighbourDuplicate(s))
}

func removeNeighbourDuplicate(s []string) []string {
	if len(s) <= 1 {
		return s
	}
	newLen := len(s)
	for i := 0; i < newLen - 1; i++ { // must be newLen - 1 here
		if s[i] == s[i + 1] { // duplicate
			copy(s[i:], s[i + 1:])
			newLen--
			i-- // notice i-- here, as there might be more than 2 neighbour duplicates
		}
	}
	return s[:newLen]
}