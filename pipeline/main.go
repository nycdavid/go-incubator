package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
			time.Sleep(250 * time.Millisecond)
		}
		close(naturals)
	}()

	go func() {
		for b := range naturals {
			squares <- b * b
		}
		close(squares)
	}()

	for {
		for x := range squares {
			fmt.Println(x)
		}
		fmt.Println("Channel closed.")
		break
	}
}
