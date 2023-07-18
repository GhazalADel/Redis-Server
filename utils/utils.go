package utils

import (
	"fmt"
	"strings"
)

func PrintCommand(buf []byte, n int) {
	bufContent := buf[:n]
	bufContentSplit := strings.Split(string(bufContent), "\n")
	fmt.Printf("Input Command : %s\n", bufContentSplit[2])
}
