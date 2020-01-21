package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println(CountWordsAndImages("http://www.baidu.com"))
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get failed", err)
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	a := make(map[string]int)
	count(&a, n)
	return a["words"], a["images"]
}

func count(m *map[string]int, n *html.Node) {
	if n.Type == html.ElementNode && string(n.Data) == "img" {
		(*m)["images"]++
	}
	if n.Type != html.ElementNode {
		(*m)["words"] += len(string(n.Data))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(m, c)
	}
}
