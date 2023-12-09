package util

func Intersect[T comparable](s1, s2 []T) (intersection []T) {
	m := make(map[T]uint8)
	for _, k := range s1 {
		m[k] |= (1 << 0)
	}
	for _, k := range s2 {
		m[k] |= (1 << 1)
	}
	for k, v := range m {
		inS1 := v&(1<<0) != 0
		inS2 := v&(1<<1) != 0
		if inS1 && inS2 {
			intersection = append(intersection, k)
		}
	}
	return
}

func Remove[T comparable](list *[]T, s T) {
	n := 0
	for _, x := range *list {
		if x != s {
			(*list)[n] = x
			n++
		}
	}
	(*list) = (*list)[:n]
}
