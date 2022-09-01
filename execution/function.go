package execution

/// Function is an overloaded Function

type Function struct {
	overloadings []PatternMatch
}

type PatternMatch struct {
	pattern  Pattern
	function Callable
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
