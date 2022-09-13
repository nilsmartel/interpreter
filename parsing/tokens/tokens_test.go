package tokens

import (
	"strings"
	"testing"
)

func template(s string) []string {
	return strings.Split(s, ",")
}

func strOfTokens(t []Token) []string {
	s := make([]string, 0, len(t))

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
		{"(print \"hello\")", "(,print, ,\"hello\",)"},
		{"(print \"\")", "(,print, ,\"\",)"},
	}

	for _, c := range cases {
		input := c[0]
		expected := template(c[1])

		tokens, err := Tokenize(input)

		if err != nil {
			t.Error("unexpected error while pasing", err)
		}

		s := strOfTokens(tokens)

		if !cmp(s, expected) {
			t.Error("\nexpected:", strings.Join(expected, ","), "\n     got:", strings.Join(s, ","))
		}

	}
}
