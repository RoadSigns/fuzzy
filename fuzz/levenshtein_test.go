package fuzz_test

import (
	"github.com/roadsigns/fuzzy-match/fuzz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevenshtein(t *testing.T) {
	assert.Equal(t, 3, fuzz.Levenshtein("kitten", "sitting"))
	assert.Equal(t, 8, fuzz.Levenshtein("pineapple", "watermelon"))
	assert.Equal(t, 0, fuzz.Levenshtein("pig", "pig"))
}
