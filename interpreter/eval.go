package interpreter

import (
	"errors"
	"fmt"
	"interpreter/ast"
	"interpreter/value"
)

func Call(env *Env, f *value.Function, args []value.Object) (value.Object, error) {
	if len(args) != len(f.Args) {
		return nil, errors.New(fmt.Sprint("called function with", len(args), "arguments, expected", len(f.Args)))
	}

	env = env.NewScope()
	for i, ident := range f.Args {
		// after calling new scope the error cant be null
		_ = env.Set(ident, args[i])
	}

	return Eval(env, f.Body)
}

func Eval(env *Env, expr ast.Expression) (value.Object, error) {

	switch expr.(type) {
	case *ast.ClassDefinition:
		defineClass(env, expr.(*ast.ClassDefinition))
	}
}

func defineClass(env *Env, def *ast.ClassDefinition) error {
	classInfo, err := value.NewClassInfo(def.Name, def.Fields)

	return nil
}
