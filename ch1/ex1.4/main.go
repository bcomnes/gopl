// “Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line
// occurs.”
// Excerpt From: Brian W. Kernighan. “The Go Programming Language (Addison-Wesley Professional
// Computing Series).” iBooks.

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileCounts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileCounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileCounts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t[ ", n, line)
			for file, count := range fileCounts[line] {
				fmt.Printf("%s:%d ", file, count)
			}
			fmt.Println("]")
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileCounts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fc, ok := fileCounts[input.Text()]
		if !ok {
			fc = make(map[string]int)
			fileCounts[input.Text()] = fc
		}
		fc[f.Name()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
