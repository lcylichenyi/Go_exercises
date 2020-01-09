package main

import (
	"fmt"
	"os"
	"bytes"
)

func main() {
	for _, arg := range os.Args[1:] {
		if len(arg) <= 3 {
			fmt.Printf("%v", arg)
		} else {
			fmt.Printf("%v", comma(arg))

		}
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	s = reverseString(s)
	for i, j := 0, 0; i < len(s); i, j = i+1, j+1 {
		if j > 0 && j % 3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return reverseString(buf.String())
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes) - 1; from < to; from, to = from + 1, to - 1 {
		runes[from], runes[to] = runes[to], runes[from]		
	}
	return string(runes)
}