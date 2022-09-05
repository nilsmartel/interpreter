package execution

import (
	"errors"
	"fmt"
	"interpreter/ast"
	"interpreter/value"
)

type concreteFunction struct {
	// optional variadic arguments are a nice thing to have
	// VarArg string
	Args []string
	Body ast.Expression
}

func NewFunction(arguments []string, body ast.Expression) (concreteFunction, error) {
	setArgs := make(map[string]bool, len(arguments))
	for _, ident := range arguments {
		if setArgs[ident] {
			return concreteFunction{}, errors.New("attempting to define multiple variables as " + ident)
		}

		setArgs[ident] = true
	}

	return concreteFunction{arguments, body}, nil
}

func (f *concreteFunction) Call(env *Env, args []value.Object) (value.Object, error) {
	if len(args) != len(f.Args) {
		return nil, errors.New(fmt.Sprint("called function with", len(args), "arguments, expected", len(f.Args)))
	}

	env = env.NewScope()
	for i, ident := range f.Args {
		// after calling new scope the error cant be null
		_ = env.Set(ident, args[i])
	}

	return Eval(env, f.Body)
}

// func (f *concreteFunction) Boolean() bool {
// 	return true
// }

// func (f *concreteFunction) Str() string {
// 	return "(fun [...] ...)"
// }

// func (f *concreteFunction) Class() string {
// 	return "Function"
// }
