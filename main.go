package main

import (
	"bufio"
	"fmt"
	"github.com/kmehran1106/Book-TheGoProgrammingLanguage/chapter1"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// For the Echo Function
	/*
		args := os.Args[1:]
		fmt.Println(chapter1.InefficientEcho(args))
		fmt.Println(chapter1.EfficientEcho(args))
	*/

	// For the Duplicate 1 Function
	/*
		counts := make(map[string]int)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if len(scanner.Text()) == 0 {
				break
			}
			counts[scanner.Text()]++
		}
		fmt.Printf(chapter1.PrintDuplicate(counts))
	*/

	// For the Duplicate 3 Function
	/*
		counts := make(map[string]int)
		for _, filename := range os.Args[1:] {
			data, err := ioutil.ReadFile(filename)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error in function: %v\n", err)
			}

			for _, line := range strings.Split(string(data), "\n") {counts[line]++}
		}
		fmt.Printf(chapter1.PrintDuplicate(counts))
	*/

}
