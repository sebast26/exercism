package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var ret Ints
	for _, e := range i {
		if filter(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var ret Ints
	for _, e := range i {
		if !filter(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var ret Lists
	for _, e := range l {
		if filter(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var ret Strings
	for _, e := range s {
		if filter(e) {
			ret = append(ret, e)
		}
	}
	return ret
}
