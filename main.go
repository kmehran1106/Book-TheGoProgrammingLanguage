package main

import (
	"fmt"
	"os"
	"github.com/kmehran1106/Book-TheGoProgrammingLanguage/chapter1"
)

func main() {
	args := os.Args[1:]
	fmt.Println(chapter1.InefficientEcho(args))
	fmt.Println(chapter1.EfficientEcho(args))
}
