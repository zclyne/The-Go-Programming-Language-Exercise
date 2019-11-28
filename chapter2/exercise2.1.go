package main

import (
	"fmt"
	"the-go-programming-language/chapter2/tempconv"
)

// 测试新增的开尔文温度及其转换
func main() {
	c := tempconv.Celsius(25)
	f := tempconv.Fahrenheit(100)
	k := tempconv.Kelvin(200)
	fmt.Println(tempconv.CToK(c))
	fmt.Println(tempconv.KToC(k))
	fmt.Println(tempconv.FToK(f))
	fmt.Println(tempconv.KToF(k))
}
