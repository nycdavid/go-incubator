package main

import (
	"fmt"
	"log"

	"github.com/nycdavid/go-incubator/links"
)

func main() {
	firstUrl := []string{"https://blog.golang.org"}
	worklist := make(chan []string)
	go func() { worklist <- firstUrl }()

	seen := make(map[string]bool)

	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

var tokens = make(chan struct{}, 10)

func crawl(url string) []string {
	// fmt.Println(url)
	tokens <- struct{}{}
	links, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return links
}

// Two arguments:
// 1. f: A function that receives a string as an argument and
//		returns a slice of strings
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
