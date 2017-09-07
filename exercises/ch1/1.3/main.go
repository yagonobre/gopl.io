package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func echo2(a []string) {
	s, sep := "", ""
	for _, arg := range a[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Fprintln(ioutil.Discard, s)
}

func echo3(a []string) {
	fmt.Fprintln(ioutil.Discard, strings.Join(a[1:], " "))
}
