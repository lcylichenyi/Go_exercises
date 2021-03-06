package main

import (
	"fmt"
)

func main() {
	a := &[]byte{'1','2','3'}
	reverse(a)
	fmt.Println(string(*a))
}

func reverse(s *[]byte) {
	for i, j := 0, len(*s) - 1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}