package main

import (
	"fmt"
	"math/rand"
)

func reverse(s *[9]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	x := [9]int{}
	for i := 0; i < 9; i++ {
		x[i] = rand.Intn(100)
	}
	fmt.Printf("Before Reverse: %v\n", x)
	reverse(&x)
	fmt.Printf("After Reverse: %v\n", x)
}