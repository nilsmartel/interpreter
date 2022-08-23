package parsing

import (
	"interpreter/ast"
	"strconv"
	"errors"
)

func Parse(tokens []Token) (ast.Expression, []Token, error) {
	if len(tokens) == 0 {
		return nil, tokens, Expected{Candidates: "<expr>"}
	}

	fst := tokens[0]
	rest := tokens[1:]
	switch fst.Tag {
	case Identifier:
		return ast.IdentLiteral{Value: fst.Span}, rest, nil
	case Int:
		value, err := strconv.ParseInt(fst.Span, 0, 64)
		if err != nil {
			return nil, nil, err
		}
		return ast.IntLiteral{Value: value}, rest, nil
	case Float:
		value, err := strconv.ParseFloat(fst.Span,  64)
		if err != nil {
			return nil, nil, err
		}
		return ast.FloatLiteral{Value: value}, rest, nil
	case String:
		value := parseString(fst.Span)
		return ast.StringLiteral{Value: value}, rest, nil

	case ParenOpen:
		return parseList(rest)

	case BracketOpen:
		return parseArray(rest)

	case EndOfInput:
		return nil, nil, errors.New("unexpected end of input")
	}

	return nil, nil, errors.New("should be unreachable")
}

// (a b c)
func parseList(tokens []Token) (ast.Expression, []Token, error) {
	// while not Paren Close: Parse
}

func parseString(raw string) string {
	// trim " " chars.
	raw = raw[1:]
	raw = raw[:len(raw)-1]

	buffer := ""

	escaped := true
	for _,c := range raw {
		if escaped {
			switch c {
			case 'n':
				buffer += "\n"
			case 't':
				buffer += "\t"
			case '\\':
				buffer += "\\"
			case 'r':
				buffer += "\r"
			default:
				buffer += string(c)
			}
			escaped = false
			continue
		}

		if c == '\\' {
			escaped = true
			continue
		}

		buffer += string(c)
	}

	return buffer
}

type Expected struct {
	Candidates string
	or         *Expected
}

func (e Expected) Error() string {
	if e.or != nil {
		return e.Candidates + " " + e.or.Error()
	}
	return e.Candidates
}

