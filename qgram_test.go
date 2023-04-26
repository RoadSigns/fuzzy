package fuzzy_test

import (
	"github.com/roadsigns/fuzzy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQGram(t *testing.T) {
	expected := 0.5
	got := fuzzy.QGram("John Jones", "Jon Jones", 3)
	assert.Equal(t, expected, got)
}

func TestQGramNegativeValue(t *testing.T) {
	expected := 0.0
	got := fuzzy.QGram("Fish", "Fish", 0)
	assert.Equal(t, expected, got)

	expected = 0.0
	got = fuzzy.QGram("Fish", "Fish", -1)
	assert.Equal(t, expected, got)
}

func BenchmarkQGram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fuzzy.QGram("John Jones", "Jon Jones", 3)
	}
}
