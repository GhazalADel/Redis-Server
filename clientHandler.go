package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func extractCommand(response []byte) string {
	// Response starts with "*1\r\n$5\r\n", so we skip the first 8 characters.
	// We look for the next "\r\n" to find the end of the command.
	end := 8
	for response[end] != '\r' || response[end+1] != '\n' {
		end++
	}
	return string(response[8 : end+2]) // Include the "\r\n" at the end of the command
}
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

		command := string(buf[:n])

		// Log the received command to the server's console
		fmt.Println("Received command:", extractCommand(buf[:n]))

		res := "Invalid Command"

		//normalize
		command = strings.TrimSpace(strings.ToLower(command))
		if command == "ping" {
			res = lc.PING()
		} else if strings.HasPrefix(command, "set ") {
			res = lc.SET(command)
		} else if strings.HasPrefix(command, "setnx ") {

		}

		_, err = conn.Write([]byte("+" + res + "\r\n"))
		if err != nil {
			fmt.Println("Error writing command: ", err.Error())
			os.Exit(1)
		}
	}
}
