package connhandler

import (
	"log"
	"net"
)

type handler func(net.Conn)

// Usage: Listen("localhost:8000")

func Listen(host string, handlerFunc handler) {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handlerFunc(conn)
	}
}
