package fuzzy_test

import (
	"github.com/roadsigns/fuzzy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJaro(t *testing.T) {
	expected := 0.8412698412698412
	got := fuzzy.Jaro("gorilla", "guerrilla")
	assert.Equal(t, expected, got)

	expected = 1
	got = fuzzy.Jaro("test", "test")
	assert.Equal(t, expected, got)
}

func BenchmarkJaro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fuzzy.Jaro("gorilla", "guerrilla")
	}
}
