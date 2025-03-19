package check

import (
	"fmt"
	"time"

	"github.com/r-happy/mulprec_go/src/mulprec"
)

func GetPi() {
	var pi mulprec.NUMBER
	mulprec.Clear(&pi)
	s := time.Now()
	NewtonNoSqrt(&pi)
	fmt.Printf("\npi: ")
	mulprec.Display(&pi)
	fmt.Println("")
	fmt.Printf("time: %f\n", time.Now().Sub(s).Seconds())
}

func NewtonNoSqrt(target *mulprec.NUMBER) {
	var denominator1, denominator2, numerator mulprec.NUMBER
	var pi, tmp_pi, zero, twoi mulprec.NUMBER
	var t1, t2 mulprec.NUMBER

	mulprec.Clear(&denominator1)
	mulprec.Clear(&denominator2)
	mulprec.Clear(&numerator)
	mulprec.Clear(&pi)
	mulprec.Clear(&tmp_pi)
	mulprec.Clear(&zero)
	mulprec.Clear(&twoi)
	mulprec.Clear(&t1)
	mulprec.Clear(&t2)
	mulprec.Clear(target)

	tmp_pi.N[mulprec.SHIFT] = 1
	pi.N[mulprec.SHIFT] = 1

	for i := 0; true; i++ {
		if i%200 == 0 {
			fmt.Print("#")
		}

		mulprec.SetInt(&twoi, int64(2*i+3))
		mulprec.SetInt(&t1, int64(2*i+1))
		mulprec.Multiple(&tmp_pi, &t1, &numerator)

		// denominator
		mulprec.SetInt(&denominator1, int64(8*(i+1)))
		mulprec.SetInt(&denominator2, int64(2*i+3))

		if mulprec.Compare(&numerator, &zero) == 0 {
			break
		}

		mulprec.Divide(&numerator, &denominator1, &t1)
		mulprec.Divide(&t1, &denominator2, &t2)

		mulprec.Add(&pi, &t2, &pi)

		mulprec.Multiple(&t2, &twoi, &tmp_pi)
	}

	mulprec.SetInt(&t1, int64(3))
	mulprec.Multiple(&t1, &pi, target)
}
