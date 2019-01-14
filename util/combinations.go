package util

// Based loosely on the python itertools implementation
func Combinations(iterable []int, r int) <-chan []int {
	ch := make(chan []int)

	go func() {
		defer close(ch)
		n := len(iterable)
		if r > n {
			// no possible combos
			return
		}

		indices := make([]int, r)
		c := make([]int, r)
		for i := range c {
			indices[i] = i
			c[i] = iterable[i]
		}
		ch <- c

		for {
			var i int
			for i = r - 1; i >= 0; i-- {
				if indices[i] != i+n-r {
					break
				}
			}
			if i == -1 {
				return
			}
			indices[i]++

			for j := i + 1; j < r; j++ {
				indices[j] = indices[j-1] + 1
			}

			c := make([]int, r)
			for k, idx := range indices {
				c[k] = iterable[idx]
			}
			ch <- c
		}
	}()

	return ch
}
