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
	return append(months, m)
}

func isNewMonth(a, b Entry) bool {
	return a.Date().Month() != b.Date().Month()
}
