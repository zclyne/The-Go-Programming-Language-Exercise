package main

import "fmt"

func main() {
	bytes := []byte("Hello world")
	fmt.Println(string(replaceSpaces(bytes)))
	bytes = []byte("Hello    world")
	fmt.Println(string(replaceSpaces(bytes)))
	bytes = []byte("HelloWorld")
	fmt.Println(string(replaceSpaces(bytes)))
	bytes = []byte("    HelloWorld")
	fmt.Println(string(replaceSpaces(bytes)))
	bytes = []byte("HelloWorld    ")
	fmt.Println(string(replaceSpaces(bytes)))
}

func replaceSpaces(bytes []byte) []byte {
	if len(bytes) <= 1 {
		return bytes
	}
	newLen := len(bytes)
	for i := 0; i < newLen - 1; i++ { // notice newLen - 1 here
		if bytes[i] == ' ' && bytes[i + 1] == bytes[i] { // neighbor space
			copy(bytes[i:], bytes[i + 1:])
			newLen--
			i-- // notice i-- here
		}
	}
	return bytes[:newLen]
}