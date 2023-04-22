package fuzzy_test

import (
	"github.com/roadsigns/fuzzy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevenshtein(t *testing.T) {
	assert.Equal(t, 3, fuzzy.Levenshtein("kitten", "sitting"))
	assert.Equal(t, 8, fuzzy.Levenshtein("pineapple", "watermelon"))
	assert.Equal(t, 0, fuzzy.Levenshtein("pig", "pig"))
}

func BenchmarkLevenshtein(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fuzzy.Levenshtein("kitten", "sitting")
	}
}
