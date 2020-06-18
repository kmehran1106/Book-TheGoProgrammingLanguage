package main

import (
	"fmt"
	"flag"
)

var str = flag.String("s", "default", "string to check adjacent uniques")

func unique(x []string) []string {
	c := 0
	for r, w := 1, 0; r < len(x); r++ {
		if x[r] != x[w] {
			w++
			x[w] = x[r]
		} else {
			c++
		}
		
	}
	return x[:len(x)-c]
}

func main() {
	flag.Parse()
	x := make([]string, len(*str))

	for i, v := range *str {
		x[i] = string(v)
	}
	fmt.Printf("Before Operation: %v\n", x)

	y := unique(x)
	
	fmt.Printf("After Operation: %v\n", y)

}