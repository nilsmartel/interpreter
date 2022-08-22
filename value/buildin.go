package value

import (
	"errors"
	"fmt"
	"interpreter/ast"
)

type NilClass struct {
}

func (i NilClass) Boolean() bool {
	return false
}

func (i NilClass) Str() string {
	return "nil"
}

var nilclassinfo = ClassInfo{name: "Nil"}

func (i NilClass) Info() *ClassInfo {
	return &nilclassinfo
}

type BoolClass struct {
	value bool
}

func (i BoolClass) Boolean() bool {
	return i.value
}

func (i BoolClass) Str() string {
	if i.value {
		return "true"
	}
	return "false"
}

var boolclassinfo = ClassInfo{name: "Bool"}

func (i BoolClass) Info() *ClassInfo {
	return &boolclassinfo
}

type IntClass struct {
	value int64
}

func (i IntClass) Boolean() bool {
	return i.value != 0
}

func (i IntClass) Str() string {
	return fmt.Sprint(i.value)
}

var intclassinfo = ClassInfo{name: "Int"}

func (i IntClass) Info() *ClassInfo {
	return &intclassinfo
}

type FloatClass struct {
	value float64
}

func (c *FloatClass) Boolean() bool {
	return c.value != 0.0
}

func (f FloatClass) Str() string {
	return fmt.Sprint(f.value)
}

var floatclassinfo = ClassInfo{name: "Float"}

func (f FloatClass) Info() *ClassInfo {
	return &floatclassinfo
}

type StringClass struct {
	value string
}

func (s *StringClass) Boolean() bool {
	return true
}

func (f *StringClass) Str() string {
	return f.value
}

var stringclassinfo = ClassInfo{name: "String"}

func (f *StringClass) Info() *ClassInfo {
	return &stringclassinfo
}

type Function struct {
	Args []string
	Body ast.Expression
}

func NewFunction(arguments []string, body ast.Expression) (Function, error) {
	setArgs := make(map[string]bool, len(arguments))
	for _, ident := range arguments {
		if setArgs[ident] == true {
			return Function{}, errors.New("attempting to define multiple variables as " + ident)
		}

		setArgs[ident] = true
	}

	return Function{arguments, body}, nil
}

func (f *Function) Boolean() bool {
	return true
}

func (f *Function) Str() string {
	return "(fun [...] ...)"
}

var functionclassinfo = ClassInfo{name: "Function"}

func (f *Function) Info() *ClassInfo {
	return &functionclassinfo
}
