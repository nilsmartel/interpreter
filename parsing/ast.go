package parsing

import (
	"errors"
	"interpreter/ast"
	"strconv"
)

func Parse(tokens []Token) (ast.Expression, []Token, error) {
	ts := make([]Token, 0)

	for _, t := range tokens {
		if t.Tag == Whitespace {
			continue
		}
		ts = append(ts, t)
	}

	return ParseCleaned(ts)
}

func ParseCleaned(tokens []Token) (ast.Expression, []Token, error) {
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
		value, err := strconv.ParseFloat(fst.Span, 64)
		if err != nil {
			return nil, nil, err
		}
		return ast.FloatLiteral{Value: value}, rest, nil
	case String:
		value := parseString(fst.Span)
		return ast.StringLiteral{Value: value}, rest, nil

	case ParenOpen:
		return parseExpr(rest)

	case BracketOpen:
		return parseArray(rest)

	case EndOfInput:
		return nil, nil, errors.New("unexpected end of input")
	}

	return nil, nil, errors.New("should be unreachable")
}

func parseArray(tokens []Token) (ast.Expression, []Token, error) {
	expr, rest, err := parseList(tokens, ParenClosing, ")")
	if err != nil {
		return nil, rest, err
	}

	return ast.ArrayLiteral{Values: expr}, rest, nil
}

// (a b c)
// or
// (<3 a b c)
func parseExpr(tokens []Token) (ast.Expression, []Token, error) {
	expr, rest, err := parseList(tokens, ParenClosing, ")")
	if err != nil {
		return nil, rest, err
	}

	if len(expr) == 0 {
		// subs with nil (e.g. do nothing)
		return ast.NilLiteral{}, rest, nil
	}

	if tokens[0].Tag == Heart {
		call, rest, err := parseExpr(tokens[1:])
		if err != nil {
			return nil, nil, err
		}

		if call, ok := call.(ast.NamedCall); ok {
			return ast.Coroutine{Call: call}, rest, nil
		}

		return nil, nil, errors.New("only named function calls are allowed to be async for now")
	}

	fst := expr[0]
	expr = expr[1:]

	if str, ok := fst.(ast.IdentLiteral); ok {
		ty := str.Value
		switch ty {
		case "do":
			return ast.DoFlow{Statements: expr}, rest, nil
		case "if":
			if len(expr) != 2 {
				return nil, rest, errors.New("expected precisely 2 arguments to if")
			}

			return ast.IfFlow{Condition: fst, True: expr[0], False: expr[1]}, rest, nil
		case "and":
			return ast.AndFlow{Arguments: expr}, rest, nil
		case "or":
			return ast.OrFlow{Arguments: expr}, rest, nil
		case "let":
			if len(expr) != 3 {
				return nil, rest, errors.New("expected precisely 3 arguments to let")
			}
			if ident, ok := expr[0].(ast.IdentLiteral); ok {
				val := expr[1]
				body := expr[2]

				return ast.VariableDefiniton{Ident: ident.Value, Value: val, Body: body}, rest, nil
			}
			return nil, rest, Expected{Candidates: "<ident>"}
		}

		return ast.NamedCall{Function: ty, Arguments: expr}, rest, nil
	}

	return ast.Call{Function: fst, Arguments: expr}, rest, nil
}

func parseList(tokens []Token, closingTag int, expectedClosing string) ([]ast.Expression, []Token, error) {
	exprs := make([]ast.Expression, 0)

	for len(tokens) > 0 && tokens[0].Tag != closingTag {
		expr, rest, err := ParseCleaned(tokens)
		if err != nil {
			return nil, nil, err
		}

		exprs = append(exprs, expr)
		tokens = rest
	}

	if len(tokens) == 0 {
		return nil, nil, Expected{Candidates: "<expr> " + expectedClosing}
	}

	return exprs, tokens[1:], nil
}

func parseString(raw string) string {
	// trim " " chars.
	raw = raw[1:]
	raw = raw[:len(raw)-1]

	buffer := ""

	escaped := true
	for _, c := range raw {
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
