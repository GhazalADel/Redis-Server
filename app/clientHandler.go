package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"redis/utils"
	"strings"
)

func HandleConnection(conn net.Conn, lc *LocalCache) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			//no more data sent
			if err == io.EOF {
				break
			} else {
				//fmt.Println("Error reading command: ", err.Error())
				//os.Exit(1)
				return
			}
		}

		command := utils.GetCommand(buf, n)
		fmt.Println(command)

		res := "Invalid Command"

		//normalize
		command = strings.TrimSpace(strings.ToLower(command))
		if command == "ping" {
			res = lc.PING()
		} else if strings.HasPrefix(command, "set ") {
			res = lc.SET(command)
		}

		_, err = conn.Write([]byte("+" + res + "\r\n"))
		if err != nil {
			fmt.Println("Error writing command: ", err.Error())
			os.Exit(1)
		}
	}
}
