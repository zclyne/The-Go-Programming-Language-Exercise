package main

import "fmt"

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)

// 这里无法使用itoa，因为1000无法用幂来表示
func main() {
	fmt.Println("KB: ", KB)
	fmt.Println("MB: ", MB)
	fmt.Println("GB: ", GB)
	fmt.Println("TB: ", TB)
	fmt.Println("PB: ", PB)
	fmt.Println("EB: ", EB)
	// 下面两句会报错constant overflows int
	// fmt.Println("ZB: ", ZB)
	// fmt.Println("YB: ", YB)
	fmt.Println("YB / ZB: ", YB / ZB)
}
