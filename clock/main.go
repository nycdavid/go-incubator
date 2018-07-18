package main

import (
	"io"
	"net"
	"os"
	"time"

	"github.com/nycdavid/go-incubator/connhandler"
)

func main() {
	connhandler.Listen("0.0.0.0:8000", handleConn)
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
