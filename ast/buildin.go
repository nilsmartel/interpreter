package ast

import "errors"

type Tag = int

const (
	Do = iota
	Async
	Await

	Add
	Subtract
	Multiply
	Modulus
	Divide
	Power

	OrChain
	AndChain

	If
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
	case "or":
		return OrChain, nil
	case "and":
		return AndChain, nil
	case "if":
		return If, nil
	}

	return 0, errors.New("not a buildin function")
}
