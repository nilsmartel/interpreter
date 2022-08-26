package main

import (
	"fmt"
	"interpreter/execution"
	"interpreter/parsing"
	"interpreter/value"
	"io/ioutil"
	"os"
)

func main() {
	filename := "sample/1-print.lisp" // os.Args[1]
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

	env := execution.NewEnv()
	// TODO populate env with defaults
	buildin(env)

	ast, rest, err := parsing.ParseCleaned(tokens)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := execution.Eval(env, ast)
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

		res, err = execution.Eval(env, ast)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Println(res.Str())
}

func buildin(env *execution.Env) {
	env.DefineGlobal("true", value.NewBool(true))
	env.DefineGlobal("false", value.NewBool(false))

	env.DefineGlobal("println", value.NewNativeFunction(func(o []value.Object) (value.Object, error) {
		for _, obj := range o {
			fmt.Print(obj.Str())
		}
		fmt.Println()

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
