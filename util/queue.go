package util

func Pop[T any](xs *[]T) (x T) {
	x, (*xs) = (*xs)[len((*xs))-1], (*xs)[:len((*xs))-1]
	return
}
