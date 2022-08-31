package execution

import (
	"errors"
	"fmt"
	"interpreter/ast"
	"interpreter/value"
)

type ClassInfo struct {
	name     string
	size     int
	fieldIds map[string]int
	methods  map[string](Function)
}

func NewClassInfo(name string, fields []string, functions []ast.FunctionDefinition) (ClassInfo, error) {
	size := len(fields)

	fieldIds := make(map[string]int, size)

	for index, ident := range fields {
		if _, ok := fieldIds[ident]; ok {
			return ClassInfo{}, errors.New("class has multiple fields with identifier " + ident)
		}

		fieldIds[ident] = index
	}

	methods := make(map[string]Function)

	for _, f := range functions {
		if _, ok := fieldIds[f.Name]; ok {
			return ClassInfo{}, errors.New(f.Name + "cant be both a field and a method on class" + name)
		}

		methods[f.Name] = Function{Args: f.Args, Body: f.Body}
	}

	return ClassInfo{name, size, fieldIds, methods}, nil
}

func (c *ClassInfo) MakeInstance(env *Env, values []value.Object) (value.Object, error) {
	if len(values) != c.size {
		return nil, errors.New(fmt.Sprint(
			"failed to create instance of class",
			c.name+".",
			"Expected", c.size, "fields as arguments. Got",
			len(values),
		))
	}

	return &Class{
		fields: values,
		info:   c,
	}, nil
}

type Class struct {
	fields []value.Object
	info   *ClassInfo
}

func (c *Class) Boolean() bool {
	return true
}

func (c *Class) Class() string {
	return c.info.name
}

func (c *Class) Method(ident string) (Function, error) {
	if fn, ok := c.info.methods[ident]; ok {
		return fn, nil
	}

	return Function{}, errors.New("no method " + ident + " on class " + c.info.name)
}

func (c *Class) Get(ident string) (value.Object, error) {
	if id, ok := c.info.fieldIds[ident]; ok {
		return c.fields[id], nil
	}

	return nil, errors.New("no field " + ident + " on class " + c.info.name)
}

// TODO this is the place to do final fields
func (c *Class) Set(ident string, value value.Object) error {
	if id, ok := c.info.fieldIds[ident]; ok {
		c.fields[id] = value
		return nil
	}

	return errors.New("no field " + ident + " on class " + c.info.name)
}

func (c *Class) Str() string {
	// e.g. (Point3d 1. 2. 3.)
	s := "(" + c.info.name

	for _, v := range c.fields {
		s += " " + v.Str()
	}

	return s + ")"
}

func (c *Class) Info() *ClassInfo {
	return c.info
}
