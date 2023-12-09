package util

type Permuter[T comparable] struct {
	slice []T
	perm  []int
}

func NewPermuter[T comparable](slice []T) *Permuter[T] {
	return &Permuter[T]{
		slice: slice,
		perm:  make([]int, len(slice)),
	}
}

func (p *Permuter[T]) NextPerm() []T {
	if p.perm[0] >= len(p.perm) {
		return nil
	}

	result := make([]T, len(p.slice))
	copy(result, p.slice)
	for i, v := range p.perm {
		result[i], result[i+v] = result[i+v], result[i]
	}

	cyclePerm(p.perm)
	return result
}

func (p *Permuter[T]) Permutations() chan []T {
	ch := make(chan []T)

	go func() {
		defer close(ch)
		for next := p.NextPerm(); next != nil; next = p.NextPerm() {
			ch <- next
		}
	}()

	return ch
}

func cyclePerm(perm []int) {
	// Cycle the permutation
	for i := len(perm) - 1; i >= 0; i-- {
		if i == 0 || perm[i] < len(perm)-i-1 {
			perm[i]++
			break
		}
		perm[i] = 0
	}
}
