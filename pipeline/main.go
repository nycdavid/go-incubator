package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}

func counter(naturals chan int) {
	for i := 0; i < 100; i++ {
		naturals <- i
		time.Sleep(250 * time.Millisecond)
	}
	close(naturals)
}

func squarer(naturals <-chan int, squares chan<- int) {
	for b := range naturals {
		squares <- b * b
	}
	close(squares)
}

func printer(squares chan int) {
	for {
		for x := range squares {
			fmt.Println(x)
		}
		fmt.Println("Channel closed.")
		break
	}
}
