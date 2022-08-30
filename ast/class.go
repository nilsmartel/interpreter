package ast

type ClassDefinition struct {
	Name    string
	Fields  []string
	Methods []FunctionDefinition
}
