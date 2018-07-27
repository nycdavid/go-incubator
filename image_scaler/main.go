package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	files := []string{
		"dog1.jpg",
		"dog2.jpg",
		"dog3.jpg",
		"dog4.jpg",
		"dog5.jpg",
		"dog6.jpg",
		"dog7.jpg",
	}

	for range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
		}()
	}

	wg.Wait()
}
