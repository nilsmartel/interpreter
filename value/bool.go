package value

type BoolClass struct {
	value bool
}

var globalTrue = BoolClass{value: true}
var globalFalse = BoolClass{value: false}

func NewBool(b bool) Object {
	if b {
		return &globalTrue
	}

	return &globalFalse
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

func (i BoolClass) Class() string {
	return "Bool"
}
