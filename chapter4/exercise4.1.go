package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	sha1, sha2 := sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X"))
	fmt.Println(sha1)
	fmt.Println(sha2)
	fmt.Println(countBit(sha1, sha2))
}

// SHA256哈希码是[32]byte数组
// 实现方法：对sha1和sha2中的每个对应位置的byte做异或操作，取出不同的bit
// 然后统计这个byte中的非0bit数
func countBit(sha1 [32]byte, sha2 [32]byte) int {
	result := 0
	for i := 0; i < 32; i++ {
		diffBits := sha1[i] ^ sha2[i]
		for diffBits > 0 {
			result += int(diffBits) & 1
			diffBits = diffBits >> 1
		}
	}
	return result
}