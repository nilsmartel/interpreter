package parsing

const (
	Identifier = iota
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

type Token struct {
	Tag  int
	Span string
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
	}

	// check if we have whitespace
	ws, rest := takeWhitespace(input)
	if ws != "" {
		return Token{Tag: Whitespace, Span: ws}, rest, nil
	}

	// must be an identifier
	ident, rest := takeNonWhitespace(input)
	return Token{Tag: Identifier, Span: ident}, rest, nil
}

func Tokenize(input string) ([]Token, error) {
	tokens := make([]Token, 0)

	for {
		t, rest, err := NextToken(input)
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, t)

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
