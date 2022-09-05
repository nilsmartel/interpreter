package execution

import (
	"errors"
	"fmt"
	"interpreter/ast"
	"interpreter/value"
)

type bytecodeFunction struct {
	// Note, these are not actually bytecode functions.
	// they are evaulated using a tree walking interpreter, which is different.
	// I had a hard time coming up with a better name.

	// optional variadic arguments are a nice thing to have
	// VarArg string
	Args []string
	Body ast.Expression
}

func NewFunction(arguments []string, body ast.Expression) (bytecodeFunction, error) {
	setArgs := make(map[string]bool, len(arguments))
	for _, ident := range arguments {
		if setArgs[ident] {
			return bytecodeFunction{}, errors.New("attempting to define multiple variables as " + ident)
		}

		setArgs[ident] = true
	}

	return bytecodeFunction{arguments, body}, nil
}

func (f *bytecodeFunction) Call(env *Env, args []value.Object) (value.Object, error) {
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
