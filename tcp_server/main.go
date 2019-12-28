package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// run telnet localhost 8080 to see the response message from the TCP server

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well I hope!")
		conn.Close()
	}
}
