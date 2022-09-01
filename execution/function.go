package execution

/// Function is an overloaded Function

type Function struct {
	overloadings []PatternMatch
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
