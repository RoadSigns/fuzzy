package fuzz_test

import (
	"github.com/roadsigns/fuzzy/fuzz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatch(t *testing.T) {
	assert.False(t, fuzz.Match("watermelon", "water"))

	assert.True(t, fuzz.Match("watermelon", "watermelon"))

	assert.False(t, fuzz.Match("applejacks", "watermelon"))

	assert.True(t, fuzz.Match("water", "watermelon"))
	assert.True(t, fuzz.Match("melon", "watermelon"))
	assert.True(t, fuzz.Match("wtrmln", "watermelon"))
	assert.True(t, fuzz.Match("wn", "watermelon"))
}

func BenchmarkMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fuzz.Match("wtrmln", "watermelon")
	}
}
