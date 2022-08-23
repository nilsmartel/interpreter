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
