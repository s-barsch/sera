package entry

import (
	"fmt"
)

type Entry struct {
	Parent *Struct
	Object interface{}
}

func (e *Entry) O() interface{} {
	return e.Object
}

func NewEntry(path string, parent *Struct) (*Entry, error) {
	return nil, fmt.Errorf("not implemented")
}
