package value

import (
	"errors"
	"fmt"
	"interpreter/ast"
)

type IntClass struct {
	value int64
}

func (i IntClass) Str() string {
	return fmt.Sprint(i.value)
}

var intclassinfo = ClassInfo{Name: "Int"}

func (i IntClass) Info() ClassInfo {
	return intclassinfo
}

type FloatClass struct {
	value float64
}

func (f FloatClass) Str() string {
	return fmt.Sprint(f.value)
}

var floatclassinfo = ClassInfo{Name: "Float"}

func (f FloatClass) Info() ClassInfo {
	return floatclassinfo
}

type StringClass struct {
	value string
}

func (f StringClass) Str() string {
	return f.value
}

var stringclassinfo = ClassInfo{Name: "String"}

func (f StringClass) Info() ClassInfo {
	return stringclassinfo
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
