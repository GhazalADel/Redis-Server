package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"redis/utils"
)

func HandleConnection(conn net.Conn) {
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

		command := utils.GetCommand(buf, n)
		fmt.Println(command)

		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("Error writing command: ", err.Error())
			os.Exit(1)
		}
	}
}
