package utils

import (
	"fmt"
	"strings"
)

func GetCommand(buf []byte, n int) string {
	bufContent := buf[:n]
	fmt.Printf("%v\t", bufContent)
	bufContentSplit := strings.Split(string(bufContent), "\n")
	for _, v := range bufContentSplit {
		fmt.Printf("%v\t", v)
	}
	return bufContentSplit[2]
}
