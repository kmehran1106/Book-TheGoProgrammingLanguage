package main

import (
	"flag"
	"crypto/sha256"
	"fmt"
	"strconv"
)

var a = flag.String("a", " ", "enter first string to compare")
var b = flag.String("b", " ", "enter second string to compare")

func popCountWithBitShift(b byte) uint8 {
	/*
		This function finds the number of bits which are set in the given byte
	*/
	var pop uint8
	for i := uint8(0); i < uint8(8); i++ {
		fmt.Printf("Before Shift: %v\t", b)
		if b&1 == 1 {
			pop++
		}
		b = b >> 1
		fmt.Printf("After Shift: %v\n", b)
	}
	return pop
}

func popCountStringMatch(b byte) uint8 {
	/*
		This function finds the number of bits which are set in the given byte
	*/
	p := strconv.FormatInt(int64(b), 2)
	var pop uint8
	for _, v := range p {
		if string(v) == "1" {
			pop++
		}
	}
	return pop
}

func main() {
	var diff int
	flag.Parse()
	x := sha256.Sum256([]byte(*a))
	y := sha256.Sum256([]byte(*b))
	
	if len(x) != len(y) {
		diff = - 1
	} else {
		for i := 0; i < len(x); i++ {
			c := x[i] ^ y[i]
			diff += int(popCountWithBitShift(c))
		}
	}
	fmt.Println(diff)
}