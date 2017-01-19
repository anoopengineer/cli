package util

import (
	"os"
	"strings"
)

//GetBinaryName returns the executable name
func GetBinaryName() string {
	binaryName := os.Args[0]
	parts := strings.Split(binaryName, "/")
	return parts[len(parts)-1]
}
