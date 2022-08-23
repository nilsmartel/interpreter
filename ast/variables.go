package ast

// (let x y body)
type VariableDefiniton struct {
	Ident string
	Value Expression
	Body  Expression
}
