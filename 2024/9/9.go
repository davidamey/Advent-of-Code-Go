package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// diskMap := util.ParseInts(string(util.MustReadFile("example")), "")
	diskMap := util.ParseInts(string(util.MustReadFile("input")), "")

	fmt.Println("p1=", p1(diskMap))
	fmt.Println("p2=", p2(diskMap))
}

func p1(diskMap []int) int {
	blocks, files, spaces := parseDiskMap(diskMap)

	fId := len(files) - 1
	for _, s := range spaces {
		if s.blockId > files[fId].blockId {
			break
		}
		for s.free() > 0 {
			f := files[fId]
			l := len(f.data) - 1
			s.data = append(s.data, f.data[l])
			f.data = f.data[:l:cap(f.data)]
			if len(f.data) == 0 {
				fId--
			}
		}
	}

	return checksum(blocks)
}

func p2(diskMap []int) int {
	blocks, files, _ := parseDiskMap(diskMap)

	// Check and move each file once (except the first)
	for fId := len(files) - 1; fId >= 1; fId-- {
		file := files[fId]
		fSize := len(file.data)
		for _, b := range blocks[:file.blockId] {
			if fSize <= b.free() { // it fits, move
				b.data = append(b.data, file.data...)
				files[fId].clear()
				break
			}
		}
	}

	return checksum(blocks)
}

func parseDiskMap(diskMap []int) (blocks, files, spaces []*block) {
	for i, b := range diskMap {
		block := newBlock(i, b)
		blocks = append(blocks, block)

		if i%2 == 0 {
			block.data = newFile(len(files), b)
			files = append(files, block)
		} else {
			spaces = append(spaces, block)
		}
	}
	return
}

func checksum(blocks []*block) (csum int) {
	i := 0
	for _, b := range blocks {
		for _, x := range b.data {
			csum += i * x
			i++
		}
		i += b.free()
	}
	return
}

type block struct {
	blockId int
	data    []int
}

func (b *block) String() string {
	var sb strings.Builder
	for i := 0; i < cap(b.data); i++ {
		if i < len(b.data) {
			sb.WriteRune(rune('0' + b.data[i]))
		} else {
			sb.WriteRune('.')
		}
	}
	return sb.String()
}

func (b *block) clear() {
	b.data = b.data[:0:cap(b.data)]
}

func (b *block) free() int {
	return cap(b.data) - len(b.data)
}

func newBlock(blockId int, size int) *block {
	data := make([]int, 0, size)
	return &block{blockId: blockId, data: data}
}

func newFile(id, size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = id
	}
	return data
}
