package main

import (
	"sync"
)

var (
	nums = []int{
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
		1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16,
	}
)

func main() {
	var wg sync.WaitGroup
	squares := make([]int, len(nums))
	ch := make(chan []int)

	for i, num := range nums {
		wg.Add(1)
		go func(i int, num int) {
			ch <- []int{i, num * num}
			wg.Done()
		}(i, num)
	}

	go func() {
		for rec := range ch {
			squares[rec[0]] = rec[1]
		}
	}()

	wg.Wait()
}

func main() {
	var squares []int
	for _, num := range nums {
		squares = append(squares, num*num)
	}
}