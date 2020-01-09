package main

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {

	m := 123456789.99

	p := message.NewPrinter(language.English)
	fmt.Println(p.Sprintf("%.2f", m))

}