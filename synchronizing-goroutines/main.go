package main

//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println("Beginning")
//
//	go func() {
//		fmt.Println("Middle")
//	}()
//
//	fmt.Println("End")
//}

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Beginning")

	wg.Add(1)
	go func() {
		fmt.Println("Middle")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("End")
}
