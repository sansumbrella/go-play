// dup prints the text of each line that appears more than once
// in the provided input files, or, if no input files are provided,
// in the standard input, preceded by its count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	sources := make(map[string]map[string]bool)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, sources)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts, sources)
		}
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, count := range counts {
		if count > 1 {
			ss, sep := "", ""
			for filename := range sources[line] {
				ss += sep + filename
				sep = ", "
			}
			fmt.Printf("%d\t%s\t%s\n", count, line, ss)
			// Calling Println avoids showing the 'D' from input termination
			// fmt.Println(strconv.Itoa(count) + ": " + line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, sources map[string]map[string]bool) {
	filename := f.Name()
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if sources[input.Text()] == nil {
			sources[input.Text()] = make(map[string]bool)
		}
		sources[input.Text()][filename] = true
	}
}
