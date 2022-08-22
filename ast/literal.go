package ast

type IdentLiteral struct {
	Value string
}

type BoolLiteral struct {
	Value bool
}

type IntLiteral struct {
	Value int64
}

type FloatLiteral struct {
	Value float64
}

// TODO might want to do this more complex, include string interpolation
type StringLiteral struct {
	Value string
}

// (fun [x y z] (+ x (+ y z)))
type LambdaLiteral struct {
	Arguments []string
	Body      Expression
}
