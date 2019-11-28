package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 从命令行参数中获取要转换的数值，如果没有参数，则从标准输入中获得
// 分别将其作为英尺和米，做相互转换
func main() {
	// 必须在这里先把num和err都声明好，如果这里不声明err，在if...else...内部用num, err :=的话
	// 会提示num未使用，无法编译
	var num float64
	var err error
	if len(os.Args) == 1 { // 无参数
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		num, err = strconv.ParseFloat(input.Text(), 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
	} else {
		num, err = strconv.ParseFloat(os.Args[1], 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("%.2fm is %.2ffeet\n", num, meterToFoot(num))
	fmt.Printf("%.2ffeet is %.2fm\n", num, footToMeter(num))
}

// 米转英尺
func meterToFoot(m float64) float64 {
	return 3.28 * m
}

// 英尺转米
func footToMeter(f float64) float64 {
	return f / 3.28
}
