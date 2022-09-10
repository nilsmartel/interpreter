package value

type Array struct {
	Values []Object
}

func NewArray(values ...Object) Object {
	return &Array{Values: values}
}

func (a *Array) Boolean() bool {
	return true
}

func (a *Array) Str() string {
	s := "["
	for _, v := range a.Values {
		s += " " + v.Str()
	}
	return s + "]"
}

func (a *Array) Class() string {
	return "Array"
}

// var arrayMethods = make(map[string]NativeFunction, 0)

// func defArrayMethods() {
// 	entry := func(ident string, method func(arr *Array, args []Object) (Object, error)) {
// 		arrayMethods["len"] = *NewNativeFunction(func(o []Object) (Object, error) {
// 			self := o[0].(*Array)
// 			return method(self, o[1:])
// 		})
// 	}

// 	entry("len", func(arr *Array, args []Object) (Object, error) {
// 		return NewInt(int64(len(arr.values))), nil
// 	})

// }

// func (a *Array) Methods(ident string) (NativeFunction, error) {
// 	if m, ok := arrayMethods[ident]; ok {
// 		return m, nil
// 	}
// 	return NativeFunction{}, noSuchMethodError(ident, "Array")
// }
