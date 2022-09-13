package execution

import (
	"errors"
	"interpreter/ast"
	"interpreter/value"
)

type bytecodeFunction struct {
	// Note, these are not actually bytecode functions.
	// they are evaulated using a tree walking interpreter, which is different.
	// I had a hard time coming up with a better name.

	// optional variadic arguments are a nice thing to have
	// VarArg string
	// Args []string
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

	return bytecodeFunction{body}, nil
}

func (f *bytecodeFunction) call(env *Env) (value.Object, error) {
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
