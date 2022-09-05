package execution

import (
	"errors"
	"interpreter/value"
	"strings"
)

/// Function is an overloaded Function

type Function struct {
	overloadings []PatternMatch
}

type PatternMatch struct {
	pattern  Pattern
	function Callable
}

type Pattern interface {
	/// matches checks if a set of values matches the pattern a method
	/// expects and, if everything's fine, returns a mapping of identifiers to these values
	matches(arguments []value.Object) (map[string]value.Object, error)
}

type FullUntyped struct {
	args []string
	// "" <-> no variadic argument
	variadic string
}

func (f *FullUntyped) matches(arguments []value.Object) (map[string]value.Object, error) {
	if len(arguments) < len(f.args) {
		remaining := f.args[len(arguments):]
		return nil, errors.New("missing arguments " + strings.Join(remaining, ", "))
	}

	idents := make(map[string]value.Object, len(f.args)+1)

	for i := range f.args {
		ident := f.args[i]
		value := arguments[i]
		idents[ident] = value
	}

	if f.variadic != "" {
		rest := arguments[:len(f.args)]

		idents[f.variadic] = value.NewArray(rest...)
	}

	return idents, nil
}

// types of patterns:

// exactly 3 arguments
// (x y z)

// anything goes (e.g. print function)
// (& x)

// at least 2 arguments
// (x y & rest)

// 2 arguments
// arg1.type == Array
// (x:Array f)

//
// ("hello")
// (0) (...)
// (1) (...)
// (n) (...)

func (f *Function) Str() string {
	return "(...)"
}

func (i *Function) Class() string {
	return "Function"
}

func (c *Function) Boolean() bool {
	return true
}
