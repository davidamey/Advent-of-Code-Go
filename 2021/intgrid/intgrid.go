package intgrid

func Parse(lines []string) (grid []int, w, h int) {
	h, w = len(lines), len(lines[0])
	grid = make([]int, w*h)
	for y, l := range lines {
		for x, r := range l {
			grid[y*w+x] = int(r - 48)
		}
	}
	return
}

func Adjacent(i, w, h int, diagonals bool) []int {
	includeLeft := i%w > 0
	includeRight := i%w < w-1
	includeTop := i >= w
	includeBottom := i < (h-1)*w

	adj := make([]int, 0, 8)
	if includeTop {
		adj = append(adj, i-w)
	}
	if includeBottom {
		adj = append(adj, i+w)
	}
	if includeLeft {
		adj = append(adj, i-1)
	}
	if includeRight {
		adj = append(adj, i+1)
	}

	if !diagonals {
		return adj
	}

	if includeTop && includeLeft {
		adj = append(adj, i-w-1)
	}
	if includeTop && includeRight {
		adj = append(adj, i-w+1)
	}
	if includeBottom && includeLeft {
		adj = append(adj, i+w-1)
	}
	if includeBottom && includeRight {
		adj = append(adj, i+w+1)
	}

	return adj
}
