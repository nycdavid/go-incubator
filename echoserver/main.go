package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/nycdavid/go-incubator/connhandler"
)

func main() {
	connhandler.Listen("0.0.0.0:8000", handleConn)
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
