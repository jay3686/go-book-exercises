// Dup prints the text of each line that appears more than once
// in STDIN, preceded by it's count. It also accepts a list of
// file names as args
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countNames := make(map[string]map[string]struct{})
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "", countNames, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, countNames, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
			for name := range countNames[line] {
				fmt.Printf("\t%s", name)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, fileName string, countNames map[string]map[string]struct{}, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if countNames[input.Text()] == nil {
			countNames[input.Text()] = make(map[string]struct{})
		}
		countNames[input.Text()][fileName] = struct{}{}
	}
	// NOTE: ignoring errors for now bc lazy
}
