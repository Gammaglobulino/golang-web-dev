package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
		go handle(conn)

	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	var i int
	var rMethod, rURI string

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			xs := strings.Fields(ln) //slice of strings from string
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD", rMethod)
			fmt.Println("URI", rURI)
		}
		if ln == "" {
			fmt.Println("End of connection")
			break
		}
		i++

	}
	body := "CHECK THE RESPONSE BODY USING A DEV TOOL OR CURL"
	body += "\n"
	body += rMethod
	body += "\n"
	body += rURI
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

}
