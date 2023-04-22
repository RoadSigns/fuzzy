package fuzz_test

import (
	"github.com/roadsigns/fuzzy/fuzz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSorensenDice(t *testing.T) {
	expect := 0.8421052631578947
	got := fuzz.SorensenDice("elon musk", "colon musk")
	assert.Equal(t, expect, got)
}

func BenchmarkSorensenDice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fuzz.SorensenDice("elon musk", "colon musk")
	}
}
