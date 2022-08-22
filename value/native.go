package value

type NativeFunction struct {
	fn func([]Object) (Object, error)
}

func NewNativeFunction(fn func([]Object) (Object, error)) *NativeFunction {
	return NewNativeFunction(fn)
}

func (f *NativeFunction) Call(values []Object) (Object, error) {
	return f.fn(values)
}

func (f *NativeFunction) Str() string {
	return ":native code:"
}

var info = ClassInfo{name: "Native Function"}

func (i *NativeFunction) Info() *ClassInfo {
	return &info
}

func (c *NativeFunction) Boolean() bool {
	return true
}
