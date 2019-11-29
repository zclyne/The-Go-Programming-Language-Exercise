package main

import "fmt"

func main() {
	fmt.Println(judge("", ""))
	fmt.Println(judge("xiaomi", "mxaoii"))
	fmt.Println(judge("xiaomi", "xiaomi"))
	fmt.Println(judge("xiaomi", "dami"))
	fmt.Println(judge("apple", "abcde"))
}

// 创建数组count用于记录各个字符出现的次数，长度设为256是默认了字符串中只包含ascii字符
// 遍历s和t，s中出现的字符记为出现次数+1
// t中出现的字符记为出现次数-1
// 若s和t含有完全相同的字符而顺序不同，最后数组count中的每个元素值都应该为0
func judge(s, t string) bool {
	// 当s和t的长度不相等时，显然不可能有同样的字符数，所以是false
	// s和t完全相等时，既有相同的字符，又有相同的顺序，不满足打乱的条件，所以是false
	if len(s) != len(t) || s == t {
		return false
	}
	var count [256]int
	for i := 0; i < len(s); i++ {
		count[byte(s[i])]++
		count[byte(t[i])]--
	}
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}