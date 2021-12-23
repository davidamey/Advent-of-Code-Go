package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	r := newReader(string(util.MustReadFile("input")))

	// r = newReader("8A004A801A8002F478")
	// r = newReader("D2FE28")
	// r = newReader("38006F45291200")
	// r = newReader("C200B40A82")

	p1, p2, _ := r.parsePacket()
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

type reader []int

func (r *reader) parsePacket() (versionTotal, value, size int) {
	initialLength := len(*r)
	defer func() {
		size = initialLength - len(*r)
	}()

	versionTotal = r.readInt(3)
	id := r.readInt(3)

	if id == 4 {
		value = r.readLiteral()
		return
	}

	var subPacketValues []int
	if lengthType := r.readBits(1)[0]; lengthType == 0 {
		for totalSize := r.readInt(15); totalSize > 0; {
			vt, v, size := r.parsePacket()
			subPacketValues = append(subPacketValues, v)
			versionTotal += vt
			totalSize -= size
		}
	} else {
		for i := r.readInt(11); i > 0; i-- {
			vt, v, _ := r.parsePacket()
			subPacketValues = append(subPacketValues, v)
			versionTotal += vt
		}
	}

	switch id {
	case 0:
		value = util.IntSum(subPacketValues...)
	case 1:
		value = util.IntProduct(subPacketValues...)
	case 2:
		value = util.MinInt(subPacketValues...)
	case 3:
		value = util.MaxInt(subPacketValues...)
	case 5:
		if subPacketValues[0] > subPacketValues[1] {
			value = 1
		}
	case 6:
		if subPacketValues[0] < subPacketValues[1] {
			value = 1
		}
	case 7:
		if subPacketValues[0] == subPacketValues[1] {
			value = 1
		}
	}

	return versionTotal, value, initialLength - len(*r)
}

func (r *reader) readBits(n int) (bits []int) {
	bits, (*r) = (*r)[:n], (*r)[n:]
	return
}

func (r *reader) readInt(n int) int {
	return toInt(r.readBits(n))
}

func (r *reader) readLiteral() int {
	var bits []int
	for {
		group := r.readBits(5)
		bits = append(bits, group[1:]...)
		if group[0] == 0 {
			break
		}
	}
	return toInt(bits)
}

func toInt(bits []int) int {
	x := 0
	for _, b := range bits {
		x <<= 1
		x += b
	}
	return x
}

func newReader(packet string) reader {
	bits := make([]int, 0, 4*len(packet))
	for _, c := range packet {
		switch c {
		case '0':
			bits = append(bits, 0, 0, 0, 0)
		case '1':
			bits = append(bits, 0, 0, 0, 1)
		case '2':
			bits = append(bits, 0, 0, 1, 0)
		case '3':
			bits = append(bits, 0, 0, 1, 1)
		case '4':
			bits = append(bits, 0, 1, 0, 0)
		case '5':
			bits = append(bits, 0, 1, 0, 1)
		case '6':
			bits = append(bits, 0, 1, 1, 0)
		case '7':
			bits = append(bits, 0, 1, 1, 1)
		case '8':
			bits = append(bits, 1, 0, 0, 0)
		case '9':
			bits = append(bits, 1, 0, 0, 1)
		case 'A':
			bits = append(bits, 1, 0, 1, 0)
		case 'B':
			bits = append(bits, 1, 0, 1, 1)
		case 'C':
			bits = append(bits, 1, 1, 0, 0)
		case 'D':
			bits = append(bits, 1, 1, 0, 1)
		case 'E':
			bits = append(bits, 1, 1, 1, 0)
		case 'F':
			bits = append(bits, 1, 1, 1, 1)
		}
	}
	return bits
}
