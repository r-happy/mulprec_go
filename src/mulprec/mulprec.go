package mulprec

import (
	"fmt"
)

const BASE int64 = 1000000000
const KETA int = 2333
const SHIFT int64 = 125

type NUMBER struct {
	N    [KETA]int64
	Sign int8
}

func Clear(s *NUMBER) {
	s.Sign = 1
	for i := range KETA {
		s.N[i] = 0
	}
}

func Display(s *NUMBER) {
	if s.Sign == 1 {
		fmt.Print("+")
	} else {
		fmt.Print("-")
	}

	zero := true

	for i := KETA - 1; i >= 0; i-- {
		if zero && s.N[i] == 0 {
			continue
		}

		zero = false

		if i == int(SHIFT) {
			fmt.Printf("%d", s.N[i])
		} else {
			fmt.Printf("%09d", s.N[i])
		}
	}

	if zero {
		fmt.Println("")
	}
}

func Copy(s *NUMBER, target *NUMBER) {
	*target = *s
}

func IsZero(s *NUMBER) bool {
	for i := range KETA {
		if s.N[i] != 0 {
			return false
		}
	}
	return true
}

func SetInt(target *NUMBER, x int64) {
	Clear(target)
	target.Sign = 1
	i := 0

	for x != 0 {
		r := x % BASE
		x /= BASE

		target.N[i] = r
		i++
	}
}

/*
s1 = s2 -> target = 0
s1 > s2 -> target = 1
s1 < s2 -> target = -1
*/
func Compare(s1 *NUMBER, s2 *NUMBER) int {
	if s1.Sign > s2.Sign {
		return 1
	}

	if s1.Sign < s2.Sign {
		return -1
	}

	r_v := 0

	for i := KETA - 1; i >= 0; i-- {
		if s1.N[i] > s2.N[i] {
			r_v = 1
			break
		}
		if s1.N[i] < s2.N[i] {
			r_v = -1
			break
		}
	}

	return r_v
}

func GetKeta(s *NUMBER) int {
	keta := 0
	for i := KETA - 1; i >= 0; i-- {
		if s.N[i] != 0 {
			keta = i + 1
			break
		}
	}
	return keta
}

func ShiftLeft(s *NUMBER, target *NUMBER, n int) {
	Clear(target)
	for i := KETA - 1; i >= 0; i-- {
		if s.N[i] != 0 {
			if (i + n) >= KETA {
				break
			}
			target.N[i+n] = s.N[i]
		}
	}
	target.N[0] = 0
	target.Sign = s.Sign
}

func ShiftRight(s *NUMBER, target *NUMBER, n int) {
	Clear(target)
	for i := range KETA - n {
		if s.N[i+n] != 0 {
			target.N[i] = s.N[i+n]
		}
	}
	target.Sign = s.Sign
}

func Add(s1 *NUMBER, s2 *NUMBER, target *NUMBER) {
	var d int64
	var e int64
	e = 0

	for i := range KETA {
		if s1.N[i] == 0 && s2.N[i] == 0 && e == 0 {
			target.N[i] = 0
			continue
		}

		d = s1.N[i] + s2.N[i] + e
		target.N[i] = d % BASE
		e = d / BASE
	}

	if e != 0 {
		Clear(target)
	}
}

func Sub(s1 *NUMBER, s2 *NUMBER, target *NUMBER) {
	var h int64 = 0

	for i := range KETA {
		s1i := s1.N[i]
		s2i := s2.N[i]
		s1i -= h

		if s1i >= s2i {
			target.N[i] = s1i - s2i
			h = 0
		}

		if s1i < s2i {
			target.N[i] = BASE + s1i - s2i
			h = 1
		}
	}

	if h != 0 {
		Clear(target)
	}
}

func Increment(s *NUMBER, target *NUMBER) {
	var one NUMBER
	SetInt(&one, 1)
	Add(s, &one, target)
}

func Decrement(s *NUMBER, target *NUMBER) {
	var one NUMBER
	SetInt(&one, 1)
	Sub(s, &one, target)
}
