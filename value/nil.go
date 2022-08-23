package value

type NilClass struct {
}

var globalNil NilClass

func Nil() Object {
	return &globalNil
}

func (i NilClass) Boolean() bool {
	return false
}

func (i NilClass) Str() string {
	return "nil"
}

func (i NilClass) Class() string {
	return "Nil"
}
