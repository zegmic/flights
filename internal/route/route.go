package route

type Route struct {
	legs     map[string][]string
	airports map[string]int
}

func (f *Route) AddLeg(from string, to string) {
	if f.legs == nil {
		f.legs = make(map[string][]string)
	}
	if f.airports == nil {
		f.airports = make(map[string]int)
	}
	f.legs[from] = append(f.legs[from], to)
	f.airports[from]++
	f.airports[to]++
}

func (f *Route) Source() string {
	for k, v := range f.airports {
		if v == 1 {
			if _, ok := f.legs[k]; ok {
				return k
			}
		}
	}
	return ""
}

func (f *Route) Destination() string {
	for k, v := range f.airports {
		if v == 1 {
			// there is no leg with destination airport as a source one
			if _, ok := f.legs[k]; !ok {
				return k
			}
		}
	}
	return ""
}

func (f *Route) IsValid() bool {
	s := f.Source()
	var v []string
	f.visit(s, &v)

	return len(f.airports) == len(v)
}

func (f *Route) visit(node string, visited *[]string) {
	v := map[string]bool{}
	v[node] = true
	*visited = append(*visited, node)
	for _, e := range f.legs[node] {
		if v[e] {
			continue
		}
		f.visit(e, visited)
	}
}
