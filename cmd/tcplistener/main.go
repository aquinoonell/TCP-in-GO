package main

import (
	"MODULE_NAME/internal/request"
	"fmt"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}

		r, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatal("error", "error", err)
		}

		fmt.Printf("request line: \n")
		fmt.Printf("- Method: %s\n", r.RequestLine.Method)
		fmt.Printf("- Targer: %s\n", r.RequestLine.RequestTarget)
		fmt.Printf("- version: %s\n", r.RequestLine.HttpVersion)
	}
}
