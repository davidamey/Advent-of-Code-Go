package util

type IntPermuter struct {
	slice []int
	perm  []int
}

type StringPermuter struct {
	slice []string
	perm  []int
}

type BytePermuter struct {
	slice []byte
	perm  []int
}

func NewIntPermuter(slice []int) *IntPermuter {
	return &IntPermuter{
		slice: slice,
		perm:  make([]int, len(slice)),
	}
}

func NewStringPermuter(slice []string) *StringPermuter {
	return &StringPermuter{
		slice: slice,
		perm:  make([]int, len(slice)),
	}
}

func NewBytePermuter(slice []byte) *BytePermuter {
	return &BytePermuter{
		slice: slice,
		perm:  make([]int, len(slice)),
	}
}

func (p *IntPermuter) NextPerm() []int {
	if p.perm[0] >= len(p.perm) {
		return nil
	}

	result := make([]int, len(p.slice))
	copy(result, p.slice)
	for i, v := range p.perm {
		result[i], result[i+v] = result[i+v], result[i]
	}

	cyclePerm(p.perm)
	return result
}

func (p *StringPermuter) NextPerm() []string {
	if p.perm[0] >= len(p.perm) {
		return nil
	}

	result := make([]string, len(p.slice))
	copy(result, p.slice)
	for i, v := range p.perm {
		result[i], result[i+v] = result[i+v], result[i]
	}

	cyclePerm(p.perm)
	return result
}

func (p *BytePermuter) NextPerm() []byte {
	if p.perm[0] >= len(p.perm) {
		return nil
	}

	result := make([]byte, len(p.slice))
	copy(result, p.slice)
	for i, v := range p.perm {
		result[i], result[i+v] = result[i+v], result[i]
	}

	cyclePerm(p.perm)
	return result
}

func (p *IntPermuter) Permutations() chan []int {
	ch := make(chan []int)

	go func() {
		defer close(ch)
		for next := p.NextPerm(); next != nil; next = p.NextPerm() {
			ch <- next
		}
	}()

	return ch
}

func (p *StringPermuter) Permutations() chan []string {
	ch := make(chan []string)

	go func() {
		defer close(ch)
		for next := p.NextPerm(); next != nil; next = p.NextPerm() {
			ch <- next
		}
	}()

	return ch
}

func (p *BytePermuter) Permutations() chan []byte {
	ch := make(chan []byte)

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
