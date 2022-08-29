package ast

type ClassDefinition struct {
	Name    string
	Fields  []string
	Methods []FunctionDefinition
}

type FunctionDefinition struct {
	Name string
	Args []string
	Body Expression
}
