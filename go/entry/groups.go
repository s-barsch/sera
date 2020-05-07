package entry

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
	return append(groups, g)
}

func isNewGroup(a, b Entry) bool {
	return a.IsBlob() != b.IsBlob()
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

func (es Entries) Months() []Entries {
	months := []Entries{}
	m := Entries{}
	for i, e := range es {
		if i > 0 && isNewMonth(e, es[i-1]) {
			months = append(months, m)
			m = Entries{e}
			continue
		}
		m = append(m, e)
	}
	months = append(months, m)
	return months
}

func isNewMonth(current, before Entry) bool {
	return current.Date().Month() != before.Date().Month()
}

