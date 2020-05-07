package entry

type Entries []Entry

func (es Entries) Reverse() Entries {
	n := Entries{}
	for i := len(es) - 1; i >= 0; i-- {
		n = append(n, es[i])
	}
	return n
}

func (es Entries) First() Entry {
	if len(es) > 0 {
		return es[0]
	}
	return nil
}


