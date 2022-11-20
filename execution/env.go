package execution

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

func NewEnv() *Env {
	return &Env{globals: make(map[string]value.Object), locals: make(map[string]value.Object)}
}

func (e *Env) NewScopeFrom(locals map[string]value.Object) *Env {
	return &Env{
		globals: e.globals,
		locals:  locals,
	}
}

func (e *Env) NewScope() *Env {
	return &Env{
		globals: e.globals,
		locals:  make(map[string]value.Object),
	}
}

func (e *Env) DefineGlobalFunction(ident string, fn Callable, pattern Pattern) error {
	if obj, ok := e.globals[ident]; ok {
		// Overload method on objects
		if callable, ok := obj.(*Function); ok {
			// TODO order of overloadings should matter
			// and be more consistent, I think
			callable.overloadings = append(callable.overloadings, PatternMatch{function: fn, pattern: pattern})
		}
		return errors.New(obj.Class() + " can't be overloaded as a function")
	}

	f := NewFunction(fn, pattern)
	e.globals[ident] = &f
	return nil
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

func (e *Env) LetIn(ident string, value value.Object, f func(e *Env) (value.Object, error)) (value.Object, error) {
	// I have a feeling that in the end, overshadowing won't work like this.
	// Because of multihreading etc.
	// well. just in case I added env as a parameter

	if v, ok := e.locals[ident]; ok {
		defer func() { e.locals[ident] = v }()
	}

	e.locals[ident] = value
	ret, err := f(e)

	// reset local varible. e.g. remove let binding from scope
	delete(e.locals, ident)
	return ret, err

}

/// Define defines a new variable in local scope
func (e *Env) SetLocal(ident string, value value.Object) error {
	if _, ok := e.locals[ident]; ok {
		return errors.New(fmt.Sprint("local variable ", ident, " already defined"))
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

	return errors.New(fmt.Sprint("attempting to assign to undefined variable ", ident))
}

func (e *Env) Get(ident string) (value.Object, error) {
	if val, ok := e.locals[ident]; ok {
		return val, nil
	}

	if val, ok := e.globals[ident]; ok {
		return val, nil
	}

	return nil, errors.New(fmt.Sprint("reading undefined variable ", ident))
}
