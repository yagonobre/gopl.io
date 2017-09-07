package main

import "testing"

var t1 = []string{"echo", "A"}
var t2 = []string{"echo", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func BenchmarkEcho2T1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		echo2(t1)
	}
}

func BenchmarkEcho3T1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		echo3(t1)
	}
}

func BenchmarkEcho2T2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		echo2(t2)
	}
}

func BenchmarkEcho3T2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		echo3(t2)
	}
}
