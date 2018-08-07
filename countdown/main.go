package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown...")
	tick := time.Tick(1 * time.Second) // time.Tick returns a channel
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	countdown := 10

	for {
		select {
		case <-tick:
			fmt.Println(countdown)
			countdown--
		case <-abort:
			countdown = 0
		}
		if countdown == 0 {
			break
		}
	}
}
