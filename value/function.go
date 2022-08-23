package value

import (
	"errors"
	"interpreter/ast"
)

type Function struct {
	Args []string
	Body ast.Expression
}

func NewFunction(arguments []string, body ast.Expression) (*Function, error) {
	setArgs := make(map[string]bool, len(arguments))
	for _, ident := range arguments {
		if setArgs[ident] == true {
			return nil, errors.New("attempting to define multiple variables as " + ident)
		}

		setArgs[ident] = true
	}

	return &Function{arguments, body}, nil
}

func (f *Function) Boolean() bool {
	return true
}

func (f *Function) Str() string {
	return "(fun [...] ...)"
}

func (f *Function) Class() string {
	return "Function"
}
