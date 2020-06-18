package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var algo = flag.Int("h", 256, "enter sha algo number")

func main() {
	flag.Parse()
	x := os.Args[len(os.Args)-1]
	if *algo == 256 {
		y := sha256.Sum256([]byte(x))
		fmt.Printf("Sha%v of %v is %x\n", *algo, x, y)
	} else if *algo == 384 {
		y := sha512.Sum384([]byte(x))
		fmt.Printf("Sha%v of %v is %x\n", *algo, x, y)
	} else if *algo == 512 {
		y := sha512.Sum512([]byte(x))
		fmt.Printf("Sha%v of %v is %x\n", *algo, x, y)
	} else {
		fmt.Println("ERROR")
	}
}