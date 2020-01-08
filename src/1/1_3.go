package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	str := []string{"adsf", "dasf", "adsfdsa", "Adsfads", "adfdasf", "adsfds"}

	start_a := time.Now()
	var b string
	b = strings.Join(str, "---")
	fmt.Println(b)
	end_a := time.Now().Sub(start_a).Nanoseconds()

	start_b := time.Now()
	var a string
	for _, v := range str {
		a += v
		a += "---"		
	}
	a = a[:len(a) - 3]
	fmt.Println(a)
	end_b := time.Now().Sub(start_b).Nanoseconds()

	fmt.Println(end_a, end_b)

}
