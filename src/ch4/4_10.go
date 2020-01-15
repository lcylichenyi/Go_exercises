package main

import (
	"fmt"
	"log"
	"os"
	"github"
	// "time"
	"sort"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)


	// var a []github.Issue
	// fmt.Printf("%T", a)

	a := timeSlice(result.Items)

	fmt.Printf("%T", a)

	// for _, v := range result.Items {
		// fmt.Printf("#%-5d %9.9s %.55s\n", result.Items[v].Number, result.Items[v].User.Login, result.Items[v].Title)
	// }
	sort.Sort(timeSlice(result.Items))
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s %v \n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}

type timeSlice []*github.Issue

func (s timeSlice) Len() int {
	return len(s)
}

func (s timeSlice) Less(i, j int) bool {
	return s[i].CreatedAt.Before(s[j].CreatedAt)
}

func (s timeSlice) Swap(i,j int) {
	s[i], s[j] = s[j], s[i]
}

