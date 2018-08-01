package main

import (
	"fmt"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func crawl(url string) []string {
	fmt.Println(url)
}
