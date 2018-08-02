package main

import (
	"fmt"
	"log"

	"github.com/nycdavid/go-incubator/links"
)

func main() {
	// breadthFirst(crawl, os.Args[1:])
	crawl("https://newyork.craigslist.org/")
}

func crawl(url string) []string {
	fmt.Println(url)
	ls, err := links.Extract(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ls)
}

// Two arguments:
// 1. f: A function that receives a string as an argument and
// 		returns a slice of strings
// 2. worklist: a slice of strings
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item))
			}
		}
	}
}
