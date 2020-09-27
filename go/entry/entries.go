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

func (es Entries) ObjectType() string {
	if len(es) == 0 {
		return "tob"
	}
	return es.First().ObjectType()
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

func (es Entries) Limit(n int) Entries {
	l := len(es)
	if l <= n {
		return es
	}
	c := 0
	nu := Entries{}
	for i := 0; c < n; i++ {
		if i + 1 >= l { return nu }

		nu = append(nu, es[i])

		if es[i].Info().Private() {
			c++
		}
	}
	return nu
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
