package value

type Object interface {
	Str() string

	// TODO might make more sense to just return TypeName
	Info() *ClassInfo
}
