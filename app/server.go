package main

import (
	"fmt"
	"net"
	"os"
)

const (
	PORT string = "6390"
	IP   string = "127.0.0.1"
)

func getAddress() string {
	return IP + ":" + PORT
}

func main() {
	listener, err := net.Listen("tcp", getAddress())
	if err != nil {
		fmt.Printf("Failed to bind to port %s", PORT)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on", getAddress())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go HandleConnection(conn)
	}
}
