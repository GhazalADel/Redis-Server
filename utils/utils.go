package utils

import (
	"strings"
)

func GetCommand(buf []byte, n int) string {
	bufContent := buf[:n]
	bufContentSplit := strings.Split(string(bufContent), "\n")
	return bufContentSplit[2]
}
