package value

import "errors"

func noSuchMethodError(method, class string) error {
	return errors.New("method " + method + " not defined on class " + class)
}
