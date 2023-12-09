package util

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func OpenExample() (*os.File, error) {
	return OpenFile("example")
}

func OpenInput() (*os.File, error) {
	return OpenFile("input")
}

func OpenFile(name string) (*os.File, error) {
	dir, _ := os.Getwd()

	// If we're in 20xx base, then look in the appropriate 'day' folder
	if path.Base(dir)[0:2] == "20" {
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

func MustReadFile(filename string) []byte {
	file, _ := OpenFile(filename)
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return b
}

func MustReadFileToLines(filename string) []string {
	file, err := OpenFile(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lines, err := ReadLines(file)
	if err != nil {
		panic(err)
	}
	return lines
}

func MustReadFileToInts(filename string) []int {
	file, err := OpenFile(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	ints, err := ReadLinesToInts(file)
	if err != nil {
		panic(err)
	}
	return ints
}

func MustReadCSInts(filename string) []int {
	raw := string(MustReadFile(filename))
	return ParseInts(raw, ",")
}

func ParseInts(s, sep string) []int {
	parts := strings.Split(s, sep)
	ints := make([]int, len(parts))
	for i, p := range parts {
		x, _ := strconv.Atoi(strings.TrimSpace(p))
		ints[i] = x
	}
	return ints
}
