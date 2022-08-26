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

func tagHeart(s string) (string, string) {
	if strings.HasPrefix(s, "<3") {
		return "<3", s[2:]
	}

	return "", s
}

func takeWhitespace(s string) (string, string) {
	for i, c := range s {
		if !unicode.IsSpace(c) {
			return s[:i], s[i:]
		}
	}

	return s, ""
}

func isSpecial(c rune) bool {
	switch c {
	case '(', ')', '[', ']', '{', '}', '"', '\'':
		return true
	}

	return false
}

func takeIdent(s string) (string, string) {
	for i, c := range s {
		if unicode.IsSpace(c) || isSpecial(c) {
			return s[:i], s[i:]
		}
	}

	return s, ""
}

func takeInt(s string) (string, string) {
	for i, c := range s {
		if !unicode.IsDigit(c) {
			return s[:i], s[i:]
		}
	}

	return s, ""
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
	for i, c := range s[1:] {
		if c == '"' && !escaped {
			// increment to respect leading "
			// and in order to contain " symbol
			i += 2
			return s[:i], s[i:]
		}

		escaped = false

		if c == '\\' {
			escaped = true
		}
	}

	return s, "" // error here. Expected "
}
