package main

import (
	"fmt"
	"tempconv"
	"os"
	"strconv"
)

func main() {
	for _, val := range os.Args[1:] {
		input, err :=  strconv.Atoi(val)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		fmt.Println(tempconv.CToK(tempconv.Celsius(float64(input))))
	}
	
}