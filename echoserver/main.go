package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	Listen("0.0.0.0:8000")
}

func Listen(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	// connection has been made
	var wg sync.WaitGroup
	defer c.Close()

	input := bufio.NewScanner(c)

	go func() {
		time.NewTicker()
	}()

	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, wg)
	}

	wg.Wait()
	fmt.Println("WaitGroup drained. Closing connection.")
}

func echo(c net.Conn, shout string, delay time.Duration, wg sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}
