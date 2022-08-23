package parsing

import (
	"strings"
	"unicode"
)

// quickly stolen from https://stackoverflow.com/a/41604514
func splitAt(s string, n int) (string, string) {
	i := 0
	for j := range s {
		if i == n {
			return s[:j], s[j:]
		}
		i++
	}
	return s, ""
}

func takeWhitespace(s string) (string, string) {
	strIndex := 0
	for i, c := range s {
		if !unicode.IsSpace(c) {
			break
		}

		strIndex += i
		continue
	}

	return s[strIndex:], s[:strIndex]
}

func takeNonWhitespace(s string) (string, string) {
	strIndex := 0
	for i, c := range s {
		if !unicode.IsSpace(c) {
			break
		}
		strIndex += i
		continue
	}

	return s[strIndex:], s[:strIndex]
}

func takeInt(s string) (string, string) {
	strIndex := 0
	for i, c := range s {
		if !unicode.IsDigit(c) {
			break
		}
		strIndex += i
		continue
	}

	return s[strIndex:], s[:strIndex]
}

func takeNumber(s string) (string, bool, string) {
	n, rest := takeInt(s)

	if !strings.HasPrefix(rest, ".") {
		return n, false, rest
	}

	n += "."
	rest = rest[1:]

	dec, rest := takeInt(rest)

	n += dec

	return n, true, rest

	// TODO implement scientific notation parsing
	/*
		if !strings.HasPrefix(rest, "e") {
			return n, true, rest
		}

		n += "e"
		rest = rest[1:]
	*/
}

func takeString(s string) (string, string) {
	if !strings.HasPrefix(s, "\"") {
		return "", s
	}

	escaped := false
	i := 0
	var c rune
	for i, c = range s[1:] {
		if c == '"' && escaped == false {
			break
		}

		escaped = false

		if c == '\\' {
			escaped = true
		}
	}

	// increment to respect leading "
	i += 1
	return s[:i], s[i:]
}
