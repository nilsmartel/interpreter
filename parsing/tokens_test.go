package parsing_test

import (
	"interpreter/parsing"
	"strings"
	"testing"
)

func template(s string) []string {
	return strings.Split(s, ",")
}

func strOfTokens(t []parsing.Token) []string {
	s := make([]string, len(t))

	for _, token := range t {
		s = append(s, token.Span)
	}

	return s
}

func cmp(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestTokenization(t *testing.T) {
	cases := [][]string{
		{"(print 7)", "(,print, ,7,)"},
		{"((a b) c)", "(,(,a, ,b,), ,c,)"},
	}

	for _, c := range cases {
		input := c[0]
		expected := template(c[1])

		tokens, err := parsing.Tokenize(input)

		if err != nil {
			t.Error("unexpected error while pasing", err)
		}

		s := strOfTokens(tokens)

		if !cmp(s, expected) {
			t.Error("expected:", strings.Join(expected, ","), "\ngot:", strings.Join(s, ","))
		}

	}
}
