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

func (i *NativeFunction) Class() string {
	return "Native Function"
}

func (c *NativeFunction) Boolean() bool {
	return true
}

// TODO Native functions to be included

/*
	Async
	Await
	Add
	Subtract
	Multiply
	Modulus
	Divide
	Power

	bitor
	bitand
	bitxor
	bitshl
	bitshr
	bitnot

	// functions to deal with io

	read "filename"
	readBytes "filename"

	write "filename" str
	writebytes "filename" []int

	print
	println
	printbytes
*/
