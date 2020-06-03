package chapter1

import "fmt"

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
