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

func (i *NativeFunction) Info() string {
	return "Native Function"
}

func (c *NativeFunction) Boolean() bool {
	return true
}
