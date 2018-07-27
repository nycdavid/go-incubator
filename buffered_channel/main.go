package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	nums := []int{0, 1, 2}
	go func() {
		for _, el := range nums {
			ch <- el
		}
	}()

	for {
		if len(ch) != cap(ch) {
			fmt.Println("Buffer not full, waiting...")
			continue
		} else {
			for i := 0; i < cap(ch); i++ {
				fmt.Println(<-ch)
			}
			break
		}
	}
}
