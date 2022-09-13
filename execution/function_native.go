package execution

import "interpreter/value"

type NativeFunction struct {
	fn func(*Env, []value.Object) (value.Object, error)
}

func NewNativeFunction(fn func(*Env, []value.Object) (value.Object, error)) *NativeFunction {
	return &NativeFunction{fn}
}

func (f *NativeFunction) call(env *Env, args []value.Object) (value.Object, error) {
	return f.fn(env, args)
}

// func (f *NativeFunction) Str() string {
// 	return "(native code)"
// }

// func (i *NativeFunction) Class() string {
// 	return "Native Function"
// }

// func (c *NativeFunction) Boolean() bool {
// 	return true
// }

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
