package parsing

import "unicode"

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
		if unicode.IsSpace(c) {
			strIndex += i
			continue
		}
	}

	return s[strIndex:], s[:strIndex]
}

func takeNonWhitespace(s string) (string, string) {
	strIndex := 0
	for i, c := range s {
		if !unicode.IsSpace(c) {
			strIndex += i
			continue
		}
	}

	return s[strIndex:], s[:strIndex]
}
