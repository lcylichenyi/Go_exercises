package main

import (
	"fmt"
)

func main() {
	fmt.Println(palindrome("爱我a", "a我爱"))
}

func palindrome(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	aa := []rune(a)
	bb := []rune(b)
	for from, to := 0, len(aa) - 1; from < to; from, to = from + 1, to - 1 {
		if aa[from] != bb[to] {
			return false
		}
	}
	return true
}