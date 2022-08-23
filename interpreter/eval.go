package interpreter

import (
	"errors"
	"fmt"
	"interpreter/ast"
	"interpreter/value"
)

func call(env *Env, caller value.Object, args []value.Object) (value.Object, error) {
	switch caller.(type) {
	case *value.Function:
		f := caller.(*value.Function)
		return callFunction(env, f, args)
	case *value.NativeFunction:
		f := caller.(*value.NativeFunction)
		return f.Call(args)
		// case *value.Class:
		// 	c := caller.(*value.Class)
		// 	f, err := c.Get("call")
		// 	if err != nil {
		// 		return nil, errors.New("class " + caller.Class() + " does not defined `call` method and is not callable.")
		// 	}
		// 	return call(env, f, args)
	}

	return nil, errors.New("value of type " + caller.Class() + " is not callable")
}

func callFunction(env *Env, f *value.Function, args []value.Object) (value.Object, error) {
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
		def := expr.(*ast.ClassDefinition)
		err := defineClass(env, def)
		if err != nil {
			return nil, err
		}
		return nil, nil

	case *ast.FunctionDefinition:
		def := expr.(*ast.FunctionDefinition)
		err := defineFunc(env, def)
		if err != nil {
			return nil, err
		}
		return nil, nil

	case *ast.DoFlow:
		statements := expr.(*ast.DoFlow).Statements
		var r value.Object
		var err error

		for _, statement := range statements {
			r, err = Eval(env, statement)
			if err != nil {
				return nil, err
			}
		}

		return r, nil

	case *ast.IfFlow:
		ifStatement := expr.(*ast.IfFlow)

		res, err := Eval(env, ifStatement.Condition)
		if err != nil {
			return nil, err
		}

		if res.Boolean() {
			return Eval(env, ifStatement.True)
		}
		return Eval(env, ifStatement.False)

	case *ast.OrFlow:
		statements := expr.(*ast.OrFlow).Arguments
		var r value.Object
		var err error

		for _, statement := range statements {
			r, err = Eval(env, statement)
			if err != nil {
				return nil, err
			}

			if r.Boolean() {
				return r, nil
			}
		}

		return r, nil

	case *ast.AndFlow:
		statements := expr.(*ast.OrFlow).Arguments
		var r value.Object
		var err error

		for _, statement := range statements {
			r, err = Eval(env, statement)
			if err != nil {
				return nil, err
			}

			if !r.Boolean() {
				return r, nil
			}
		}

		return r, nil

	case *ast.NamedCall:
		c := expr.(*ast.NamedCall)

		function, err := env.Get(c.Function)
		if err != nil {
			return nil, err
		}

		args := make([]value.Object, len(c.Arguments))
		for _, arg := range c.Arguments {
			value, err := Eval(env, arg)
			if err != nil {
				return nil, err
			}

			args = append(args, value)
		}

		return call(env, function, args)

	case *ast.Call:
		c := expr.(*ast.Call)

		function, err := Eval(env, c.Function)
		if err != nil {
			return nil, err
		}

		args := make([]value.Object, len(c.Arguments))
		for _, arg := range c.Arguments {
			value, err := Eval(env, arg)
			if err != nil {
				return nil, err
			}

			args = append(args, value)
		}

		return call(env, function, args)

	// Literals
	case *ast.IdentLiteral:
		ident := expr.(*ast.IdentLiteral).Value
		return env.Get(ident)

	case *ast.BoolLiteral:
		val := expr.(*ast.BoolLiteral).Value
		return value.NewBool(val), nil

	case *ast.NilLiteral:
		return value.Nil(), nil

	case *ast.IntLiteral:
		val := expr.(*ast.IntLiteral).Value
		return value.NewInt(val), nil

	case *ast.FloatLiteral:
		val := expr.(*ast.FloatLiteral).Value
		return value.NewFloat(val), nil

	case *ast.StringLiteral:
		val := expr.(*ast.StringLiteral).Value
		return value.NewString(val), nil

	case *ast.LambdaLiteral:
		lit := expr.(*ast.LambdaLiteral)
		return value.NewFunction(lit.Arguments, lit.Body)

	case *ast.ArrayLiteral:
		expressions := expr.(*ast.ArrayLiteral).Values
		values := make([]value.Object, len(expressions))
		for _, expr := range expressions {
			v, err := Eval(env, expr)
			if err != nil {
				return nil, err
			}

			values = append(values, v)
		}
		return value.NewArray(values...), nil

	// (let x y body)
	case *ast.VariableDefiniton:
		def := expr.(*ast.VariableDefiniton)

		// first evaluate variable assignment
		val, err := Eval(env, def.Value)
		if err != nil {
			return nil, err
		}

		// overshadow what has been
		return env.LetIn(def.Ident, val, func(env *Env) (value.Object, error) {
			return Eval(env, def.Body)
		})
	}

	return nil, errors.New("unknown expression encountered")
}

func defineClass(env *Env, def *ast.ClassDefinition) error {
	classInfo, err := value.NewClassInfo(def.Name, def.Fields)
	if err != nil {
		return err
	}

	constructor := value.NewNativeFunction(classInfo.MakeInstance)

	env.DefineGlobal(def.Name, constructor)
	return nil
}

func defineFunc(env *Env, def *ast.FunctionDefinition) error {
	// TODO this is the spot to include information about codeposition in file, row, col
	// for stacktraces etc.
	function, err := value.NewFunction(def.Args, def.Body)
	if err != nil {
		return err
	}
	env.DefineGlobal(def.Name, function)
	return nil
}
