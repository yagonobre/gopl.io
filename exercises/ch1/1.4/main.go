// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Line struct {
	count map[string]int
	files map[string][]string
}

func main() {
	l := Line{
		count: make(map[string]int),
		files: make(map[string][]string),
	}

	files := os.Args[1:]
	for _, arg := range files {
		err := countLines(arg, l)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
	}
	for line, n := range l.count {
		if n > 1 {
			fmt.Printf("Count: %d\tValue: %s\tFiles:%s\n", n, line, strings.Join(l.files[line], " "))
		}
	}
}

func countLines(name string, line Line) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		t := input.Text()
		line.count[t]++
		line.files[t] = append(line.files[t], name)
	}
	// NOTE: ignoring potential errors from input.Err()

	return nil
}

//!-
