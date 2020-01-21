// ../fetch/fetch.exe http:www.baidu.com | go run ./5.2.go
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	a := make(map[string]int)
	visit(&a, doc)
	fmt.Println(a)
}

func visit(m *map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		(*m)[string(n.Data)]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(m, c)
	}
}
