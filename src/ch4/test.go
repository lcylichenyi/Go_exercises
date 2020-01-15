package main

import (
	"fmt"
)

func main() {
	a := fmt.Errorf("search query failed: %s", 404)
	fmt.Println(a)
}