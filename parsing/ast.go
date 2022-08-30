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

	return parse(ts)
}

func parse(tokens []Token) (ast.Expression, []Token, error) {
	if len(tokens) == 0 {
		return nil, tokens, Expected{Candidates: "<expr>"}
	}

	fst := tokens[0]
	rest := tokens[1:]
	switch fst.Tag {
	case ParenOpen:
		return parseBody(rest)

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

	case BracketOpen:
		return parseArray(rest)

	case EndOfInput:
		return nil, nil, errors.New("unexpected end of input")
	}

	return nil, nil, errors.New("should be unreachable. got")
}

/// Parses (x y z)
/// Class and function definitions
func parseBody(tokens []Token) (ast.Expression, []Token, error) {
	if len(tokens) < 2 {
		return nil, nil, errors.New("unexpected end of input")
	}

	// (class)
	if tokens[0].Tag == Class {
		name, tokens, err := expect(tokens[1:], Identifier, "<class name>")
		if err != nil {
			return nil, nil, err
		}
		_, tokens, err = expect(tokens, ParenOpen, "(")
		if err != nil {
			return nil, nil, err
		}

		fields := make([]string, 0)
		for i, t := range tokens {
			if t.Tag == Identifier {
				fields = append(fields, t.Span)
				continue
			}
			if t.Tag == ParenClosing {
				// advance token list
				tokens = tokens[i+1:]
				break
			}
			// illegal token encountered
			return nil, nil, Expected{Candidates: "<ident> )"}
		}

		methods := make([]ast.FunctionDefinition, 0)
		for tokens[0].Tag != ParenClosing {
			_, tokens, err = expect(tokens, Fun, "fun")
			if err != nil {
				return nil, nil, Expected{Candidates: ")", or: err.(*Expected)}
			}

			name, tokens, err = expect(tokens, Identifier, "<ident>")
			if err != nil {
				return nil, nil, err
			}

			_, tokens, err = expect(tokens, ParenOpen, "<ident>")
			if err != nil {
				return nil, nil, err
			}

			args := make([]string, 0)
			for i, t := range tokens {
				if t.Tag == Identifier {
					args = append(args, t.Span)
					continue
				}
				if t.Tag == ParenClosing {
					tokens = tokens[i+1:]
					break
				}
				return nil, nil, Expected{Candidates: "<ident> )"}
			}

			_, tokens, err := expect(tokens, ParenClosing, ")")
			if err != nil {
				return nil, nil, err
			}
			_, tokens, err = expect(tokens, ParenOpen, "(")
			if err != nil {
				return nil, nil, err
			}
			body, tokens, err := parseExpr(tokens)
			if err != nil {
				return nil, nil, err
			}

			fd := ast.FunctionDefinition{Name: name.Span, Args: args, Body: body}
			methods = append(methods, fd)
		}

		_, tokens, err = expect(tokens, ParenClosing, ")")
		if err != nil {
			return nil, nil, err
		}

		return ast.ClassDefinition{Name: name.Span, Fields: fields, Methods: methods}, tokens, nil
	}

	// (fun)
	if tokens[0].Tag == Fun && tokens[1].Tag == Identifier {
		funcName := tokens[1].Span

		_, tokens, err := expect(tokens, ParenOpen, "(")
		if err != nil {
			return nil, nil, err
		}
		// TODO ident ) ( parseExpr
		args := make([]string, 0)
		for tokens[0].Tag != Identifier {
			args = append(args, tokens[0].Span)
			tokens = tokens[1:]

			if len(tokens) == 0 {
				return nil, nil, Expected{Candidates: "<argument> )"}
			}
		}

		_, tokens, err = expect(tokens, ParenClosing, ")")
		if err != nil {
			return nil, nil, err
		}
		_, tokens, err = expect(tokens, ParenOpen, "(")
		if err != nil {
			return nil, nil, err
		}

		body, rest, err := parseExpr(tokens)
		if err != nil {
			return nil, nil, err
		}

		return ast.FunctionDefinition{Name: funcName, Args: args, Body: body}, rest, nil
	}

	return parseExpr(tokens)
}

func expect(tokens []Token, tag int, expected string) (Token, []Token, error) {
	if len(tokens) == 0 ||
		tokens[0].Tag != tag {
		return Token{}, nil, Expected{Candidates: expected}
	}

	return tokens[0], tokens[1:], nil
}

func parseArray(tokens []Token) (ast.Expression, []Token, error) {
	expr, rest, err := parseList(tokens, BracketClosing, "]")
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
		return nil, nil, errors.New("() is not allowed") // ast.NilLiteral{}, rest, nil
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
		expr, rest, err := parse(tokens)
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
