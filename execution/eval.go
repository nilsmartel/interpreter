package execution

import (
	"errors"
	"fmt"
	"interpreter/ast"
	"interpreter/value"
)

func call(env *Env, caller value.Object, args []value.Object) (value.Object, error) {
	switch caller := caller.(type) {
	case Callable:
		return caller.call(env, args)
	}

	return nil, errors.New("value of type " + caller.Class() + " is not callable")
}

func Eval(env *Env, expr ast.Expression) (value.Object, error) {
	switch expr := expr.(type) {
	case *ast.ClassDefinition:
		err := defineClass(env, expr)
		if err != nil {
			return nil, err
		}
		return nil, nil

	case *ast.FunctionDefinition:
		err := defineFunc(env, expr)
		if err != nil {
			return nil, err
		}
		return nil, nil

	case ast.DoFlow:
		var r value.Object
		var err error

		for _, statement := range expr.Statements {
			r, err = Eval(env, statement)
			if err != nil {
				return nil, err
			}
		}

		return r, nil

	case ast.IfFlow:

		res, err := Eval(env, expr.Condition)
		if err != nil {
			return nil, err
		}

		if res.Boolean() {
			return Eval(env, expr.True)
		}
		return Eval(env, expr.False)

	case ast.OrFlow:
		var r value.Object
		var err error

		for _, statement := range expr.Arguments {
			r, err = Eval(env, statement)
			if err != nil {
				return nil, err
			}

			if r.Boolean() {
				return r, nil
			}
		}

		return r, nil

	case ast.AndFlow:
		var r value.Object
		var err error

		for _, statement := range expr.Arguments {
			r, err = Eval(env, statement)
			if err != nil {
				return nil, err
			}

			if !r.Boolean() {
				return r, nil
			}
		}

		return r, nil

	case ast.NamedCall:
		return namedCall(env, &expr)

	case ast.Call:
		return callExpr(env, expr)

	// Literals
	case ast.IdentLiteral:
		return env.Get(expr.Value)

	case ast.BoolLiteral:
		return value.NewBool(expr.Value), nil

	case ast.NilLiteral:
		return value.Nil(), nil

	case ast.IntLiteral:
		return value.NewInt(expr.Value), nil

	case ast.FloatLiteral:
		return value.NewFloat(expr.Value), nil

	case ast.StringLiteral:
		return value.NewString(expr.Value), nil

	case ast.LambdaLiteral:
		// TODO we actually really want to capture the env
		// at this point
		f, err := NewBytecodeFunction(expr.Arguments, expr.Body)
		return &f, err

	case ast.ArrayLiteral:
		values := make([]value.Object, 0, len(expr.Values))
		for _, expr := range expr.Values {
			v, err := Eval(env, expr)
			if err != nil {
				return nil, err
			}

			values = append(values, v)
		}
		return value.NewArray(values...), nil

	// (let x y body)
	case ast.VariableDefiniton:
		// first evaluate variable assignment
		val, err := Eval(env, expr.Value)
		if err != nil {
			return nil, err
		}

		// overshadow what has been
		return env.LetIn(expr.Ident, val, func(env *Env) (value.Object, error) {
			return Eval(env, expr.Body)
		})
	}

	return nil, errors.New("unknown expression encountered: " + fmt.Sprintf("%+v", expr))
}

func callExpr(env *Env, expr ast.Call) (value.Object, error) {
	function, err := Eval(env, expr.Function)
	if err != nil {
		return nil, err
	}

	args := make([]value.Object, 0, len(expr.Arguments))
	for _, arg := range expr.Arguments {
		value, err := Eval(env, arg)
		if err != nil {
			return nil, err
		}

		args = append(args, value)
	}

	return call(env, function, args)
}

func namedCall(env *Env, expr *ast.NamedCall) (value.Object, error) {
	// Evaluate arguments
	args := make([]value.Object, 0, len(expr.Arguments))
	for _, arg := range expr.Arguments {
		value, err := Eval(env, arg)
		if err != nil {
			return nil, err
		}

		args = append(args, value)
	}

	function, err := env.Get(expr.Function)
	if err != nil {
		return nil, err
	}

	return call(env, function, args)
}

func defineClass(env *Env, def *ast.ClassDefinition) error {
	classInfo, err := NewClassInfo(def.Name, def.Fields, def.Methods)
	if err != nil {
		return err
	}

	// The constructor of a class is a native function,
	// that gets designed only once.
	constructor := NewNativeFunction(classInfo.MakeInstance)

	env.DefineGlobalFunction(def.Name, constructor, &FullUntyped{args: def.Fields})
	return nil
}

func defineFunc(env *Env, def *ast.FunctionDefinition) error {
	// TODO this is the spot to include information about codeposition in file, row, col
	// for stacktraces etc.
	bytecodeFunc, err := NewBytecodeFunction(def.Args, def.Body)
	if err != nil {
		return err
	}

	env.DefineGlobalFunction(def.Name, &bytecodeFunc, &FullUntyped{args: def.Args})
	return nil
}
