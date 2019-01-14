package main

import "fmt"

func main() {
	// examples
	// for _, s := range [][]byte{
	// 	[]byte("hijklmmn"),
	// 	[]byte("abbceffg"),
	// 	[]byte("abbcegjk"),
	// } {
	// 	vc, _ := TestValidChars(s)
	// 	fmt.Println(string(s), TestIncreasingStraight(s), vc, TestPairs(s))
	// }

	orig := []byte("vzbxkghb")
	p1 := NextPass(orig)
	p2 := NextPass(p1)
	fmt.Println("p1=", string(p1))
	fmt.Println("p2=", string(p2))
}

func NextPass(current []byte) []byte {
	newPass := make([]byte, len(current))
	copy(newPass, current)

	CycleString(newPass)
	for {
		if valid, pos := TestValidChars(newPass); !valid {
			newPass[pos]++
			for i := pos + 1; i < len(newPass); i++ {
				newPass[i] = 'a'
			}
			continue
		}

		if !TestIncreasingStraight(newPass) || !TestPairs(newPass) {
			CycleString(newPass)
			continue
		}

		break
	}
	return newPass
}

func TestIncreasingStraight(s []byte) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i+2]-s[i] == 2 && s[i+1]-s[i] == 1 {
			return true
		}
	}
	return false
}

func TestValidChars(s []byte) (bool, int) {
	for i, b := range s {
		if b == 'i' || b == 'o' || b == 'l' {
			return false, i
		}
	}
	return true, -1
}

func TestPairs(s []byte) bool {
	pairs := make(map[byte]int)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			pairs[s[i]]++
			i++
		}
	}
	return len(pairs) >= 2
}

func CycleString(b []byte) {
	for i := len(b) - 1; i > 0; i-- {
		if b[i] == 122 {
			b[i] = 97
			continue
		}
		b[i]++
		break
	}
}
