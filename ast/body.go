package ast

type FunctionDefinition struct {
	Name string
	Args []string
	Body Expression
}
