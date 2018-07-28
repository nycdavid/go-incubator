package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/nycdavid/go-incubator/connhandler"
)

func main() {
	connhandler.Listen("0.0.0.0:8000", handleConn)
}

func handleConn(c net.Conn) {
	// connection has been made
	var wg sync.WaitGroup
	defer c.Close()
	input := bufio.NewScanner(c)
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
