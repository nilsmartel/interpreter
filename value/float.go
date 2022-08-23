package value

import "fmt"

type FloatClass struct {
	value float64
}

func NewFloat(value float64) Object {
	return &FloatClass{value}
}

func (c *FloatClass) Boolean() bool {
	return c.value != 0.0
}

func (f FloatClass) Str() string {
	return fmt.Sprint(f.value)
}

func (f FloatClass) Class() string {
	return "Float"
}
