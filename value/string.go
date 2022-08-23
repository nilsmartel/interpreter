package value

type StringClass struct {
	value string
}

func NewString(value string) Object {
	return &StringClass{value}
}

func (s *StringClass) Boolean() bool {
	return true
}

func (f *StringClass) Str() string {
	return f.value
}

func (f *StringClass) Class() string {
	return "String"
}
