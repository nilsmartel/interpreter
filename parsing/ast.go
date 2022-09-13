package parsing

import (
	"errors"
	"interpreter/ast"
	"interpreter/parsing/tokens"
	"strconv"
)

func Parse(ts []tokens.Token) (ast.Expression, []tokens.Token, error) {
	cleaned := make([]tokens.Token, 0)

	for _, t := range ts {
		if t.Tag == tokens.Whitespace {
			continue
		}
		cleaned = append(cleaned, t)
	}

	return parse(cleaned)
}

func parse(ts []tokens.Token) (ast.Expression, []tokens.Token, error) {
	if len(ts) == 0 {
		return nil, ts, Expected{Candidates: "<expr>"}
	}

	fst := ts[0]
	rest := ts[1:]
	switch fst.Tag {
	case tokens.ParenOpen:
		return parseBody(rest)

	case tokens.Identifier:
		return ast.IdentLiteral{Value: fst.Span}, rest, nil
	case tokens.Bool:
		value := fst.Span == "true"
		return ast.BoolLiteral{Value: value}, rest, nil
	case tokens.Int:
		value, err := strconv.ParseInt(fst.Span, 0, 64)
		if err != nil {
			return nil, nil, err
		}
		return ast.IntLiteral{Value: value}, rest, nil
	case tokens.Float:
		value, err := strconv.ParseFloat(fst.Span, 64)
		if err != nil {
			return nil, nil, err
		}
		return ast.FloatLiteral{Value: value}, rest, nil
	case tokens.String:
		value := parseString(fst.Span)
		return ast.StringLiteral{Value: value}, rest, nil

	case tokens.BracketOpen:
		return parseArray(rest)

	case tokens.EndOfInput:
		return nil, nil, errors.New("unexpected end of input")
	}

	return nil, nil, errors.New("should be unreachable. got")
}

func parseClass(ts []tokens.Token) (ast.Expression, []tokens.Token, error) {

	name, ts, err := expect(ts, tokens.Identifier, "<class name>")
	if err != nil {
		return nil, nil, err
	}
	_, ts, err = expect(ts, tokens.ParenOpen, "(")
	if err != nil {
		return nil, nil, err
	}

	fields := make([]string, 0)
	for i, t := range ts {
		if t.Tag == tokens.Identifier {
			fields = append(fields, t.Span)
			continue
		}
		if t.Tag == tokens.ParenClosing {
			// advance token list
			ts = ts[i+1:]
			break
		}
		// illegal token encountered
		return nil, nil, Expected{Candidates: "<ident> )"}
	}

	methods := make([]ast.FunctionDefinition, 0)
	for ts[0].Tag != tokens.ParenClosing {
		_, ts, err = expect(ts, tokens.Fun, "fun")
		if err != nil {
			return nil, nil, Expected{Candidates: ")", or: err.(*Expected)}
		}

		name, ts, err = expect(ts, tokens.Identifier, "<ident>")
		if err != nil {
			return nil, nil, err
		}

		_, ts, err = expect(ts, tokens.ParenOpen, "<ident>")
		if err != nil {
			return nil, nil, err
		}

		args := make([]string, 0)
		for i, t := range ts {
			if t.Tag == tokens.Identifier {
				args = append(args, t.Span)
				continue
			}
			if t.Tag == tokens.ParenClosing {
				ts = ts[i+1:]
				break
			}
			return nil, nil, Expected{Candidates: "<ident> )"}
		}

		_, ts, err = expect(ts, tokens.ParenClosing, ")")
		if err != nil {
			return nil, nil, err
		}
		_, ts, err = expect(ts, tokens.ParenOpen, "(")
		if err != nil {
			return nil, nil, err
		}
		body, rest, err := parseExpr(ts)
		ts = rest
		if err != nil {
			return nil, nil, err
		}

		fd := ast.FunctionDefinition{Name: name.Span, Args: args, Body: body}
		methods = append(methods, fd)
	}

	_, ts, err = expect(ts, tokens.ParenClosing, ")")
	if err != nil {
		return nil, nil, err
	}

	return ast.ClassDefinition{Name: name.Span, Fields: fields, Methods: methods}, ts, nil
}

func parseNamedFunction(ts []tokens.Token) (ast.Expression, []tokens.Token, error) {

	// fun <ident>
	if !(ts[0].Tag == tokens.Fun && ts[1].Tag == tokens.Identifier) {
		return nil, nil, Expected{Candidates: "fun_<ident>"}
	}
	funcName := ts[1].Span

	// (args) TODO move to own function and pattern match
	_, ts, err := expect(ts, tokens.ParenOpen, "(")
	if err != nil {
		return nil, nil, err
	}

	args := make([]string, 0)
	for ts[0].Tag != tokens.Identifier {
		args = append(args, ts[0].Span)
		ts = ts[1:]

		if len(ts) == 0 {
			return nil, nil, Expected{Candidates: "<argument> )"}
		}
	}

	_, ts, err = expect(ts, tokens.ParenClosing, ")")
	if err != nil {
		return nil, nil, err
	}

	// TODO this is wrong. We can also return constants!
	_, ts, err = expect(ts, tokens.ParenOpen, "(")
	if err != nil {
		return nil, nil, err
	}

	body, rest, err := parseExpr(ts)
	if err != nil {
		return nil, nil, err
	}

	return ast.FunctionDefinition{Name: funcName, Args: args, Body: body}, rest, nil
}

/// Parses (x y z)
/// Class and function definitions
func parseBody(ts []tokens.Token) (ast.Expression, []tokens.Token, error) {
	if len(ts) < 2 {
		return nil, nil, errors.New("unexpected end of input")
	}

	// (class)
	if ts[0].Tag == tokens.Class {
		return parseClass(ts[1:])
	}

	// (fun <ident>)
	if f, r, err := parseNamedFunction(ts); err == nil {
		return f, r, err
	}

	return parseExpr(ts)
}

func expect(ts []tokens.Token, tag int, expected string) (tokens.Token, []tokens.Token, error) {
	if len(ts) == 0 ||
		ts[0].Tag != tag {
		return tokens.Token{}, nil, Expected{Candidates: expected}
	}

	return ts[0], ts[1:], nil
}

func parseArray(ts []tokens.Token) (ast.Expression, []tokens.Token, error) {
	expr, rest, err := parseList(ts, tokens.BracketClosing, "]")
	if err != nil {
		return nil, rest, err
	}

	return ast.ArrayLiteral{Values: expr}, rest, nil
}

// (a b c)
// or
// (<3 a b c)
func parseExpr(ts []tokens.Token) (ast.Expression, []tokens.Token, error) {
	expr, rest, err := parseList(ts, tokens.ParenClosing, ")")
	if err != nil {
		return nil, rest, err
	}

	if len(expr) == 0 {
		// subs with nil (e.g. do nothing)
		return nil, nil, errors.New("() is not allowed") // ast.NilLiteral{}, rest, nil
	}

	if ts[0].Tag == tokens.Heart {
		call, rest, err := parseExpr(ts[1:])
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
			if len(expr) != 3 {
				return nil, rest, errors.New("expected precisely 3 arguments to if")
			}

			return ast.IfFlow{Condition: expr[0], True: expr[1], False: expr[2]}, rest, nil
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

func parseList(ts []tokens.Token, closingTag int, expectedClosing string) ([]ast.Expression, []tokens.Token, error) {
	exprs := make([]ast.Expression, 0)

	for len(ts) > 0 && ts[0].Tag != closingTag {
		expr, rest, err := parse(ts)
		if err != nil {
			return nil, nil, err
		}

		exprs = append(exprs, expr)
		ts = rest
	}

	if len(ts) == 0 {
		return nil, nil, Expected{Candidates: "<expr> " + expectedClosing}
	}

	return exprs, ts[1:], nil
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
