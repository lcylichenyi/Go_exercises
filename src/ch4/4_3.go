package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(&[]int{1,2,3}))
}

func reverse(s *[]int) *[]int{
	for i, j := 0, len(*s) - 1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
	return s
}