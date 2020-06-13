package main

import (
	"fmt"
)

func main() {
	var p *int
	fmt.Println(p)

	var x int = 1
	p = &x
	fmt.Println(p)
	fmt.Println(*p)
}