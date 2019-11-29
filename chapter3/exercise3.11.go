package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma2("123456"))
	fmt.Println(comma2("1234.56"))
	fmt.Println(comma2("-123456"))
	fmt.Println(comma2("-1234.56"))
	fmt.Println(comma2("-1234.56789"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma2(s string) string {
	// 处理第一个字符可能为正负号的情况
	sign := ""
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if s[0] == '-' {
			sign = "-"
		}
	}
	// 获得小数点位置，从而判断是否为浮点数
	dotIndex := strings.Index(s, ".")
	if dotIndex != -1 { // is float
		return sign + commaHelper(s[:dotIndex]) + "." + s[dotIndex + 1:] // 小数部分不做加逗号处理
	} else { // is integer
		return sign + commaHelper(s)
	}
}

func commaHelper(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaHelper(s[:n-3]) + "," + s[n-3:]
}