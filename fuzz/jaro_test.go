package fuzz_test

import (
	"github.com/roadsigns/fuzzy-match/fuzz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJaro(t *testing.T) {
	expected := 0.8412698412698412
	got := fuzz.Jaro("gorilla", "guerrilla")
	assert.Equal(t, expected, got)

	expected = 1
	got = fuzz.Jaro("test", "test")
	assert.Equal(t, expected, got)
}

func BenchmarkJaro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fuzz.Jaro("gorilla", "guerrilla")
	}
}
