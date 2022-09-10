package execution

import (
	"errors"
	"interpreter/value"
	"strings"
)

type Pattern interface {
	/// matches checks if a set of values matches the pattern a method
	/// expects and, if everything's fine, returns a mapping of identifiers to these values
	matches(arguments []value.Object) (map[string]value.Object, error)
}

type FullUntyped struct {
	args []string
	// "" <-> no variadic argument
	variadic string
}

func (f *FullUntyped) matches(arguments []value.Object) (map[string]value.Object, error) {
	if len(arguments) < len(f.args) {
		remaining := f.args[len(arguments):]
		return nil, errors.New("missing arguments " + strings.Join(remaining, ", "))
	}

	idents := make(map[string]value.Object, len(f.args)+1)

	for i := range f.args {
		ident := f.args[i]
		value := arguments[i]
		idents[ident] = value
	}

	if f.variadic != "" {
		rest := arguments[len(f.args):]

		idents[f.variadic] = value.NewArray(rest...)
	}

	return idents, nil
}

type Arguments struct {
	args     []Argument
	variadic string
}

func (a *Arguments) matches(values []value.Object) (map[string]value.Object, error) {
	if len(values) < len(a.args) {
		return nil, errors.New("not enough arguments")
	}

	m := make(map[string]value.Object)

	for i, arg := range a.args {
		value := values[i]
		err := arg.match(value, &m)
		if err != nil {
			return nil, err
		}
	}

	if a.variadic != "" {
		rest := values[len(a.args):]
		m[a.variadic] = value.NewArray(rest...)
	}

	return m, nil
}

type Argument interface {
	match(value.Object, *map[string]value.Object) error
}

type ArrayDestructure struct {
	patterns Arguments
}

func (a ArrayDestructure) match(val value.Object, idents *map[string]value.Object) error {
	if val, ok := val.(*value.Array); ok {
		values := val.Values

		m, err := a.patterns.matches(values)
		if err != nil {
			return err
		}

		// apply all map keys
		for k, v := range m {
			(*idents)[k] = v
		}

		return nil
	}

	return errors.New("expected Array, got" + val.Class())
}

type Constant struct {
	value value.Object
}

func (c Constant) match(value value.Object, idents *map[string]value.Object) error {
	if !constEqual(c.value, value) {
		return errors.New(value.Str() + " does not match " + c.value.Str())
	}

	return nil
}

type Variable struct {
	identifier string
}

func (v Variable) match(value value.Object, idents *map[string]value.Object) error {
	(*idents)[v.identifier] = value

	return nil
}

type VariableBinding struct {
	identifier string
	binding    Arguments
}

func (v VariableBinding) match(value value.Object, idents *map[string]value.Object) error {
	(*idents)[v.identifier] = value

	return nil
}

type VariableTyped struct {
	identifier string
	kind       string
}

func (v VariableTyped) match(value value.Object, idents *map[string]value.Object) error {
	if v.kind != value.Class() {
		return errors.New("wrong type. got " + v.kind + " expected " + value.Class())
	}
	(*idents)[v.identifier] = value

	return nil
}

func constEqual(a value.Object, b value.Object) bool {
	// TODO this will 1 != 1.0
	// should we consider this a bug?

	switch a := a.(type) {
	case *value.NilClass:
		if _, ok := b.(*value.NilClass); ok {
			return true
		}
		return false
	case *value.BoolClass:
		if value, ok := b.(*value.BoolClass); ok {
			return a == value
		}
		return false
	case *value.IntClass:
		if value, ok := b.(*value.IntClass); ok {
			return a == value
		}
		return false
	case *value.FloatClass:
		if value, ok := b.(*value.FloatClass); ok {
			return a == value
		}
		return false
	case *value.StringClass:
		if value, ok := b.(*value.StringClass); ok {
			return a == value
		}
		return false
	}

	panic("can't compare consts like that")
}

// types of patterns:

// exactly 3 arguments
// (x y z)

// anything goes (e.g. print function)
// (& x)

// at least 2 arguments
// (x y & rest)

// 2 arguments
// arg1.type == Array
// (x:Array f)

//
// ("hello")
// (0) (...)
// (1) (...)
// (n) (...)
