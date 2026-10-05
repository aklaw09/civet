package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:6380", "address to listen on")
	flag.Parse()

	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal("starting server:", err)
	}
	log.Println("listening on", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accepting connection:", err)
			continue
		}

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("client connected:", conn.RemoteAddr())
	_, err := conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		log.Println("writing response:", err)
	}
}
