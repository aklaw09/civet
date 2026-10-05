package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting server on :", 6730)
	port := 6730
	addr := flag.String("addr", "localhost", "Address to listen on")
	flag.Parse()
	listener, err := net.Listen("tcp", *addr+":"+fmt.Sprint(port))
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())
	_, err := conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		fmt.Println("Error writing to connection:", err)
	}
}
