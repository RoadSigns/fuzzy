package fuzz_test

import (
	"github.com/roadsigns/fuzzy-match/fuzz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoundex(t *testing.T) {
	expect := "S530"
	assert.Equal(t, expect, fuzz.Soundex("Smith"))
	assert.Equal(t, expect, fuzz.Soundex("Smythe"))
}

func BenchmarkSoundex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fuzz.Soundex("Watermelon")
	}
}
