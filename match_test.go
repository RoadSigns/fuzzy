package fuzzy_test

import (
	"github.com/roadsigns/fuzzy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatch(t *testing.T) {
	assert.False(t, fuzzy.Match("watermelon", "water"))

	assert.True(t, fuzzy.Match("watermelon", "watermelon"))

	assert.False(t, fuzzy.Match("applejacks", "watermelon"))

	assert.True(t, fuzzy.Match("water", "watermelon"))
	assert.True(t, fuzzy.Match("melon", "watermelon"))
	assert.True(t, fuzzy.Match("wtrmln", "watermelon"))
	assert.True(t, fuzzy.Match("wn", "watermelon"))
}

func BenchmarkMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fuzzy.Match("wtrmln", "watermelon")
	}
}
