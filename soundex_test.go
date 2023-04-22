package fuzzy_test

import (
	"github.com/roadsigns/fuzzy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoundex(t *testing.T) {
	expect := "S530"
	assert.Equal(t, expect, fuzzy.Soundex("Smith"))
	assert.Equal(t, expect, fuzzy.Soundex("Smythe"))
}

func BenchmarkSoundex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fuzzy.Soundex("Watermelon")
	}
}
