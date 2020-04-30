package entry

type Entry struct {
	Parent *Struct
	Object interface{}
}

func (e *Entry) O() interface{} {
	return e.Object
}

