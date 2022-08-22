package value

import "errors"

type ClassInfo struct {
	name     string
	size     int
	fieldIds map[string]int
}

func NewClassInfo(name string, fields []string) (ClassInfo, error) {
	size := len(fields)

	fieldIds := make(map[string]int, size)

	for index, ident := range fields {
		if _, ok := fieldIds[ident]; ok {
			return ClassInfo{}, errors.New("class has multiple fields with identifier " + ident)
		}

		fieldIds[ident] = index
	}

	return ClassInfo{name, size, fieldIds}, nil
}

type Class struct {
	fields []Object
	info   *ClassInfo
}

func (c *Class) Get(ident string) (Object, error) {
	if id, ok := c.info.fieldIds[ident]; ok {
		return c.fields[id], nil
	}

	return nil, errors.New("no field " + ident + " on class " + c.info.name)
}

// TODO this is the place to do final fields
func (c *Class) Set(ident string, value Object) error {
	if id, ok := c.info.fieldIds[ident]; ok {
		c.fields[id] = value
		return nil
	}

	return errors.New("no field " + ident + " on class " + c.info.name)
}

func (c *Class) Str() string {
	// e.g. (Point3d 1. 2. 3.)
	s := "(" + c.info.name

	for _, v := range c.fields {
		s += " " + v.Str()
	}

	return s + ")"
}

func (c *Class) Info() *ClassInfo {
	return c.info
}
