package parsing

import (
	"interpreter/ast"
)

func Parse(tokens []Token) ([]ast.Expression, error) {

}

// (a b c)
func ParseList(tokens []Token) (ast.Expression, error) {
	if tokens[0].Tag != ParenOpen {
		return nil, Expected{Candidates: "("}
	}

	// TODO parse list of expressions
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
