package fuzzy

import (
	"unicode/utf8"
)

// Match checks to see if the two strings are a partial match and if so return true,
// if they aren't then it returns a false. Match is case-sensitive.
func Match(source, target string) bool {
	return match(source, target)
}

func match(source, target string) bool {
	sLen := len(source)
	tLen := len(target)

	if sLen > tLen {
		return false
	}

	if sLen == tLen {
		return source == target
	}

Outer:
	for _, s := range source {
		for i, t := range target {
			if t == s {
				target = target[i+utf8.RuneLen(t):]
				continue Outer
			}
		}
		return false
	}

	return true
}
