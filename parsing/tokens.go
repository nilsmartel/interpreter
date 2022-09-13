package parsing

import "fmt"

const (
	Identifier = iota

	Bool
	Int
	Float
	String

	In

	Dot

	TypeColon

	Heart
	Class
	Fun

	ParenOpen
	ParenClosing
	BracketOpen
	BracketClosing
	Whitespace

	EndOfInput

	// not yet used
	CurlyOpen
	CurlyClosing
)

func tagToStr(tag int) string {
	switch tag {
	case Identifier:
		return "Identifier"
	case Int:
		return "Int"
	case Float:
		return "Float"
	case String:
		return "String"
	case Heart:
		return "Heart"
	case Class:
		return "Class"
	case Fun:
		return "Fun"
	case ParenOpen:
		return "ParenOpen"
	case ParenClosing:
		return "ParenClosing"
	case BracketOpen:
		return "BracketOpen"
	case BracketClosing:
		return "BracketClosing"
	case Whitespace:
		return "Whitespace"
	case EndOfInput:
		return "EndOfInput"
	case CurlyOpen:
		return "CurlyOpen"
	case CurlyClosing:
		return "CurlyClosing"
	}

	return "undefined"
}

type Token struct {
	Tag  int
	Span string
}

func (t Token) String() string {
	return fmt.Sprintf("<%s> [%s]", tagToStr(t.Tag), t.Span)
}

func NextToken(input string) (Token, string, error) {
	if input == "" {
		return Token{Tag: EndOfInput}, "", nil
	}

	fst, rest := splitAt(input, 1)
	switch fst {
	case "(":
		return Token{Tag: ParenOpen, Span: fst}, rest, nil
	case ")":
		return Token{Tag: ParenClosing, Span: fst}, rest, nil
	case "[":
		return Token{Tag: BracketOpen, Span: fst}, rest, nil
	case "]":
		return Token{Tag: BracketClosing, Span: fst}, rest, nil
	case "{":
		return Token{Tag: CurlyOpen, Span: fst}, rest, nil
	case "}":
		return Token{Tag: CurlyClosing, Span: fst}, rest, nil
	case ".":
		return Token{Tag: Dot, Span: fst}, rest, nil
	case ":":
		return Token{Tag: TypeColon, Span: fst}, rest, nil
	}

	// check if we have whitespace
	if ws, rest := takeWhitespace(input); ws != "" {
		return Token{Tag: Whitespace, Span: ws}, rest, nil
	}

	if num, floating, rest := takeNumber(input); num != "" {
		tag := Int
		if floating {
			tag = Float
		}

		return Token{Tag: tag, Span: num}, rest, nil
	}

	if str, rest := takeString(input); str != "" {
		return Token{Tag: String, Span: str}, rest, nil
	}

	// must be an identifier
	ident, rest := takeIdent(input)
	tag := Identifier
	switch ident {
	case "class":
		tag = Class
	case "fun":
		tag = Fun
	case "<3":
		tag = Heart
	case "in":
		tag = In
	case "true":
		tag = Bool
	case "false":
		tag = Bool
	}
	return Token{Tag: tag, Span: ident}, rest, nil
}

func Tokenize(input string) ([]Token, error) {
	tokens := make([]Token, 0)

	for {
		t, rest, err := NextToken(input)
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, t)
		input = rest

		if rest == "" {
			break
		}
	}

	return tokens, nil
}

func YieldTokens(input string, ch chan Token) error {
	for {
		t, rest, err := NextToken(input)
		if err != nil {
			return err
		}

		ch <- t

		if rest == "" {
			break
		}
	}

	return nil
}
