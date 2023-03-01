package main

import (
	"math/rand"
	"testing"
)

var testSample = make([]int, 2000)

func init() {
	for i := 0; i < 2000; i++ {
		testSample[i] = randInt(-9, 9)
	}
}

func BenchmarkCountDumb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximumCountDumb(testSample)
	}
}

func BenchmarkCountSmart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximumCountNoBranches(testSample)
	}
}

func randInt(lower, upper int) int {
	rng := upper - lower
	return rand.Intn(rng) + lower
}
