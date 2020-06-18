package main

import (
	"fmt"
	"flag"
	"math/rand"
)

var rotate = flag.Int("r", 2, "amount of numbers to rotate. must be less than 10")

func main() {
	flag.Parse()

	x := make([]int, 0)
	for i := 0; i < 9; i++ {
		x = append(x, rand.Intn(100))
	}
	fmt.Printf("Before Rotate: %v\n", x)

	x = append(x[*rotate:], x[:*rotate]...)
	
	fmt.Printf("After Rotate: %v\n", x)

}