package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	log.Print("Starting...")
	tz, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now()
	t2 := t.In(tz)
	fmt.Println(t2.String())
}
