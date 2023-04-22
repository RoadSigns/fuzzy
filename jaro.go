package fuzzy

import "math"

func Jaro(source, target string) float64 {
	sLen := len(source)
	tLen := len(target)

	window := int(math.Floor(math.Max(float64(sLen), float64(tLen))/2.0)) - 1

	m := 0
	t := 0
	prevMatchIdx := -1

	for i := 0; i < sLen; i++ {
		start := int(math.Max(float64(0), float64(i-window)))
		end := int(math.Min(float64(tLen-1), float64(i+window)))
		for j := start; j <= end; j++ {
			if source[i] == target[j] {
				m++
				if prevMatchIdx != -1 && j < prevMatchIdx {
					t++
				}
				prevMatchIdx = j
				break
			}
		}
	}

	if m == 0 {
		return 0.0
	}
	return (1.0 / 3.0) * (float64(m)/float64(sLen) + float64(m)/float64(tLen) + (float64(m)-float64(t))/float64(m))
}
