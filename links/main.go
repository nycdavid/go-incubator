package main

import (
	"fmt"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

//func crawl(url string) []string {
//	fmt.Println(url)
//}

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
			}
		}
	}
}
