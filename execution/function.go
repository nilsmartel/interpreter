package execution

import (
	"errors"
	"interpreter/value"
)

/// Function is an overloaded Function

type Function struct {
	overloadings []PatternMatch
}

func (f *Function) call(env *Env, args []value.Object) (value.Object, error) {
	for _, fun := range f.overloadings {
		m, err := fun.pattern.matches(args)
		if err != nil {
			continue
		}

		fun.function.call(env.NewScopeFrom(m), args)
	}

	// TODO return errors from pattern matching for clearer debugging
	return nil, errors.New("No overload matching call")
}

type PatternMatch struct {
	pattern  Pattern
	function Callable
}

func (f *Function) Str() string {
	return "(...)"
}

func (i *Function) Class() string {
	return "Function"
}

func (c *Function) Boolean() bool {
	return true
}
