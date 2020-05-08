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

func (es Entries) Offset(start, end int) Entries {
	l := len(es)
	if l < start {
		return Entries{}
	}
	if end > l || end <= 0 {
		return es[start:]
	}
	return es[start:end]
}

func (es Entries) Exclude() Entries {
	l := Entries{}
	for _, e := range es {
		if e.Info()["hidden"] == "true" || e.Info()["exclude"] == "true" {
			continue
		}
		l = append(l, e)
	}
	return l
}
