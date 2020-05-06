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

func (es Entries) Groups() []Entries {
	groups := []Entries{}

	g := Entries{}

	for i, e := range es {
		if i > 0 && isNewGroup(es[i-1], e) {
			groups = append(groups, g)
			g = Entries{e}
			continue
		}
		g = append(g, e)
	}
	return  append(groups, g)
}

func isNewGroup(a, b Entry) bool {
	return a.IsBlob() != b.IsBlob()
}
