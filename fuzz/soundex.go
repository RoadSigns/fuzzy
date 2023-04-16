package fuzz

import (
	"bytes"
	"strings"
)

func Soundex(word string) string {
	word = strings.ToUpper(word)
	var code bytes.Buffer
	code.WriteByte(word[0])

	for i := 1; i < len(word); i++ {
		switch word[i] {
		case 'B', 'F', 'P', 'V':
			code.WriteByte('1')
		case 'C', 'G', 'J', 'K', 'Q', 'S', 'X', 'Z':
			code.WriteByte('2')
		case 'D', 'T':
			code.WriteByte('3')
		case 'L':
			code.WriteByte('4')
		case 'M', 'N':
			code.WriteByte('5')
		case 'R':
			code.WriteByte('6')
		default:
			// Ignore non-alphabetic characters
		}
	}

	for i := 1; i < code.Len(); i++ {
		if code.Bytes()[i] == code.Bytes()[i-1] {
			code.Truncate(i)
			i--
		}
	}

	for code.Len() < 4 {
		code.WriteByte('0')
	}

	return code.String()[:4]
}
