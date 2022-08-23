package value

type Array struct {
	values []Object
}

func NewArray(values ...Object) Object {
	return &Array{values}
}

func (a *Array) Boolean() bool {
	return true
}

func (a *Array) Str() string {
	s := "["
	for _, v := range a.values {
		s += " " + v.Str()
	}
	return s + "]"
}

func (a *Array) Class() string {
	return "Array"
}
