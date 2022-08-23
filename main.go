package main

import (
	"fmt"
	"interpreter/interpreter"
	"interpreter/parsing"
	"interpreter/value"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading file", err)
		os.Exit(1)
	}

	input := string(data)

	tokens, err := parsing.Tokenize(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	env := interpreter.NewEnv()
	// TODO populate env with defaults
	buildin(env)

	ast, rest, err := parsing.Parse(tokens)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := interpreter.Eval(env, ast)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for len(rest) > 0 {
		ast, rest, err = parsing.Parse(tokens)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		res, err = interpreter.Eval(env, ast)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Println(res.Str())
}

func buildin(env *interpreter.Env) {
	env.DefineGlobal("true", value.NewBool(true))
	env.DefineGlobal("false", value.NewBool(false))

	env.DefineGlobal("println", value.NewNativeFunction(func(o []value.Object) (value.Object, error) {
		for _, obj := range o {
			fmt.Println(obj.Str())
		}
		return value.Nil(), nil
	}))

	env.DefineGlobal("str", value.NewNativeFunction(func(o []value.Object) (value.Object, error) {
		s := ""
		for _, obj := range o {
			s += obj.Str()
		}
		return value.NewString(s), nil
	}))
}
