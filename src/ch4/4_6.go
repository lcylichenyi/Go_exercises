package main

import (
	"unicode"
	"fmt"
	"os"
)

func main() {
	for _, v := range os.Args[1:] {
		fmt.Println(cut(v))
	}
}

func cut(s string) string{
	a := []rune(s)
	for i, flag := 0, false; i < len(a); i++ {
		if unicode.IsSpace(a[i]) {
			if flag {
				a = append(a[:i], a[i+1:]...)
				i--
			} 
			flag = true
			continue
		}
		flag = false
	}
	return string(a)
} 