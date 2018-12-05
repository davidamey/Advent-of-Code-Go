package util

import (
	"bufio"
	"io"
	"os"
	"path"
	"strconv"
)

func OpenExample() (*os.File, error) {
	return openFile("example")
}

func OpenInput() (*os.File, error) {
	return openFile("input")
}

func openFile(name string) (*os.File, error) {
	dir, _ := os.Getwd()

	// If we're in 2018 base, then look in the appropriate 'day' folder
	if path.Base(dir) == "2018" {
		dir = path.Join(dir, path.Base(os.Args[0]))
	}

	return os.Open(path.Join(dir, name))
}

func ReadLinesToInts(r io.Reader) ([]int, error) {
	var result []int

	lines, err := ReadLines(r)
	if err != nil {
		return result, err
	}

	for _, l := range lines {
		x, err := strconv.Atoi(l)
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}

	return result, nil
}

func ReadLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}
