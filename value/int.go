package value

import "fmt"

type IntClass struct {
	value int64
}

func NewInt(value int64) Object {
	return &IntClass{value}
}

func (i IntClass) Boolean() bool {
	return i.value != 0
}

func (i IntClass) Str() string {
	return fmt.Sprint(i.value)
}

func (i IntClass) Class() string {
	return "Int"
}
