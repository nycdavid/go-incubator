package main

import (
	"fmt"
	"log"

	"github.com/nycdavid/go-incubator/links"
)

func main() {
	breadthFirst(crawl, []string{"https://www.google.com"})
}

func crawl(url string) []string {
	fmt.Println(url)
	links, err := links.Extract(url)
	if err != nil {
		log.Fatal(err)
	}
	return links
}

// Two arguments:
// 1. f: A function that receives a string as an argument and
// 		returns a slice of strings
// 2. worklist: a slice of strings
func breadthFirst(crawlFunc func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, crawlFunc(item)...)
			}
			fmt.Println(len(worklist))
		}
	}
}
