package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatal(err)
	}
	tzS := os.Getenv("TZ")
	fmt.Println(tzS)
	tz, err := time.LoadLocation(tzS)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue // breaks execution
		}
		go handleConn(conn, tz)
	}
}

func handleConn(c net.Conn, loc *time.Location) {
	defer c.Close()
	for {
		t := time.Now()
		t2 := t.In(loc)
		_, err := io.WriteString(c, t2.Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
