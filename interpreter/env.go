package interpreter

import (
	"errors"
	"fmt"
	"interpreter/value"
)

type Env struct {
	// Note: this is the point to declare stuff as constant.
	// e.g. make
	// globals map[string](value.Object, mutable bool)
	globals map[string]value.Object
	locals  map[string]value.Object
}

func (e *Env) NewScope() *Env {
	return &Env{
		globals: e.globals,
		locals:  make(map[string]value.Object),
	}
}

func (e *Env) DefineGlobal(ident string, value value.Object) error {
	if _, ok := e.globals[ident]; ok {
		return errors.New(fmt.Sprint(ident, "already defined"))
	}

	e.globals[ident] = value
	return nil
}

func (e *Env) SetGlobal(ident string, value value.Object) error {
	if _, ok := e.globals[ident]; !ok {
		return errors.New(fmt.Sprint("attempting to assign to undefined variable", ident))
	}

	e.globals[ident] = value
	return nil
}

/// Define defines a new variable in local scope
func (e *Env) Let(ident string, value value.Object) error {
	if _, ok := e.locals[ident]; ok {
		return errors.New(fmt.Sprint("local variable", ident, "already defined"))
	}

	e.locals[ident] = value
	return nil
}

func (e *Env) Set(ident string, value value.Object) error {
	if _, ok := e.locals[ident]; ok {
		e.locals[ident] = value
		return nil
	}

	if _, ok := e.globals[ident]; ok {
		e.globals[ident] = value
		return nil
	}

	return errors.New(fmt.Sprint("attempting to assign to undefined variable", ident))
}

func (e *Env) Get(ident string) (value.Object, error) {
	if val, ok := e.locals[ident]; ok {
		return val, nil
	}

	if val, ok := e.globals[ident]; ok {
		return val, nil
	}

	return nil, errors.New(fmt.Sprint("reading undefined variable", ident))
}
