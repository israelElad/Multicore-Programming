// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	lineToFileName := make(map[string][]string)

	if len(files) == 0 {
		countLines(os.Stdin, counts, nil)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineToFileName)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			fmt.Printf("%v\n", lineToFileName[line])
		}
	}
}
func countLines(f *os.File, counts map[string]int, lineToFileName map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !checkIfValueInSlice(lineToFileName[input.Text()], f.Name()) {
			lineToFileName[input.Text()] = append(lineToFileName[input.Text()], f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func checkIfValueInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
