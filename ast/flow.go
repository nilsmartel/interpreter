package ast

type DoFlow struct {
	Statements []Expression
}

type IfFlow struct {
	Condition Expression
	True      Expression
	False     Expression
}

type OrFlow struct {
	Arguments []Expression
}

type AndFlow struct {
	Arguments []Expression
}

/*
type Tag = int

const (
	Do = iota
	If
	OrChain
	AndChain

	// TODO rewrite as native functions
	Async
	Await
	Add
	Subtract
	Multiply
	Modulus
	Divide
	Power
)

// (do a b c d e f g)
type BuildIn struct {
	Kind      Tag
	Arguments Expression
}

func ToBuildIn(s string) (Tag, error) {
	switch s {
	case "do":
		return Do, nil
	case "or":
		return OrChain, nil
	case "and":
		return AndChain, nil
	case "if":
		return If, nil
	// TODO the rest are actually native functions
	case "<3":
		return Async, nil
	case "..":
		return Await, nil
	case "+":
		return Add, nil
	case "-":
		return Subtract, nil
	case "*":
		return Multiply, nil
	case "%":
		return Modulus, nil
	case "/":
		return Divide, nil
	case "**":
		return Power, nil
	}

	return 0, errors.New("not a buildin function")
}
*/
