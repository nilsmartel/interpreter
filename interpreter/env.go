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
	if set, ok := e.globals[ident]; ok && set != nil {
		return errors.New(fmt.Sprint(ident, "already defined"))
	}

	e.globals[ident] = value
	return nil
}

func (e *Env) SetGlobal(ident string, value value.Object) error {
	if set, ok := e.globals[ident]; !(ok && set != nil) {
		return errors.New(fmt.Sprint("attempting to assign to undefined variable", ident))
	}

	e.globals[ident] = value
	return nil
}

func (e *Env) LetIn(ident string, value value.Object, f func(e *Env) (value.Object, error)) (value.Object, error) {
	v := e.locals[ident]
	e.locals[ident] = value
	// I have a feeling that in the end, overshadowing won't work like this.
	// Because of multihreading etc.
	// well. just in case I added env as a parameter
	ret, err := f(e)

	// reset local varible. e.g. remove let binding from scope
	e.locals[ident] = v
	return ret, err

}

/// Define defines a new variable in local scope
func (e *Env) SetLocal(ident string, value value.Object) error {
	if set, ok := e.locals[ident]; ok && set != nil {
		return errors.New(fmt.Sprint("local variable", ident, "already defined"))
	}

	e.locals[ident] = value
	return nil
}

func (e *Env) Set(ident string, value value.Object) error {
	if set, ok := e.locals[ident]; ok && set != nil {
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
	if val, ok := e.locals[ident]; ok && val != nil {
		return val, nil
	}

	if val, ok := e.globals[ident]; ok && val != nil {
		return val, nil
	}

	return nil, errors.New(fmt.Sprint("reading undefined variable", ident))
}
