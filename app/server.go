package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const (
	PORT string = "6390"
	IP   string = "127.0.0.1"
)

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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			//no more data sent
			if err == io.EOF {
				break
			} else {
				fmt.Println("Error reading command: ", err.Error())
				os.Exit(1)
			}
		}

		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("Error writing command: ", err.Error())
			os.Exit(1)
		}
	}
}

func getAddress() string {
	return IP + ":" + PORT
}
