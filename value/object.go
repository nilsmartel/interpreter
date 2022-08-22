package value

type Object interface {
	Str() string

	Boolean() bool

	// TODO might make more sense to just return TypeName
	Info() *ClassInfo
}
