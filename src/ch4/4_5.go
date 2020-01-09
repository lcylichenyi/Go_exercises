package main

import (
	"fmt"
)

func main() {
	a := &[]string{"aa","aa","bb","cc", "cc"}
	clear(a)
	fmt.Println(a)
}

func clear(s *[]string) {
	for i := 0; i < len(*s) - 1; i++ {
		if (*s)[i] == (*s)[i+1] {
			// copy((*s)[i:], (*s)[i+1:])
			*s = append((*s)[0: i+1], (*s)[i+2:]...)
			i--
		}
	}
}