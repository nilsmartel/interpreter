package parsing

import (
	"testing"
)

func TestTakeString(t *testing.T) {
	s, rest := takeString("\"hello\"lots of love")

	if s != "\"hello\"" {
		t.Error("expected hello string, got", s)
	}

	if rest != "lots of love" {
		t.Error("expected 'lots of love' string, got", rest)
	}

}

func TestTakeWhitespace(t *testing.T) {
	s, rest := takeWhitespace("   \n()")

	if s != "   \n" {
		t.Error("expected whitespace")
	}

	if rest != "()" {
		t.Error("expected rest, got", rest)
	}

}
