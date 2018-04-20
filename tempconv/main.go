package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	msg := fmt.Sprintf("Sleeping for %v...", *period)
	fmt.Println(msg)
	time.Sleep(*period)
	fmt.Println("Done sleeping")
}

type MyFlag struct {
}

func (mf *MyFlag) String() string {
	return ""
}

func (mf *MyFlag) Set(arg string) error {
	return nil
}
