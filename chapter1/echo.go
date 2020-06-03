package main

import (
	"fmt"
	"os"
	"strings"
)

// EfficientEcho : X
func EfficientEcho(elements []string) string {
	joinedArgs := strings.Join(elements, " ")
	return joinedArgs
}

// InefficientEcho : Z
func InefficientEcho(elements []string) string {
	var s, sep string
	for index, arg := range elements {
		s += sep + arg
		if index == 0 {
			sep = " "
		}
	}
	return s
}

func main() {
	args := os.Args[1:]
	fmt.Println(InefficientEcho(args))
	fmt.Println(EfficientEcho(args))
}
