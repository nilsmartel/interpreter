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
