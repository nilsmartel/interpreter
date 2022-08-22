package ast

type ClassDefinition struct {
	Name   string
	Fields []string
}

type FunctionDefinition struct {
	Name string
	Args []string
	Body Expression
}
