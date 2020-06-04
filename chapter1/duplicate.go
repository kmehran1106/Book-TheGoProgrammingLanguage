package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// PrintDuplicate : Used to find Duplicate from a Map
func PrintDuplicate(counts map[string]int) string {
	returnString := ""
	for line, n := range counts {
		if n > 1 {
			returnString += fmt.Sprintf("There are %d instances of the line %s \n", n, line)
		}
	}
	return returnString
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if len(scanner.Text()) == 0 {
				break
			}
			counts[scanner.Text()]++
		}
	} else {
		for _, filename := range os.Args[1:] {
			data, err := ioutil.ReadFile(filename)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error in function: %v\n", err)
			}

			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}
		}
	}
	fmt.Printf(PrintDuplicate(counts))
}
