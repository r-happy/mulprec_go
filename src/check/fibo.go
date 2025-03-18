package check

import (
	"fmt"
	"time"

	"github.com/r-happy/mulprec_go/src/mulprec"
)

// const KETA int = 2333 この値に設定しなおす

func Fibonacci() {
	fmax := 100000
	var f0 mulprec.NUMBER
	var f1 mulprec.NUMBER
	var tmp mulprec.NUMBER

	mulprec.SetInt(&f0, 0)
	mulprec.SetInt(&f1, 1)
	mulprec.Clear(&tmp)

	s := time.Now()

	for i := 2; i <= fmax; i++ {
		mulprec.Add(&f0, &f1, &tmp)
		mulprec.Copy(&f1, &f0)
		mulprec.Copy(&tmp, &f1)
	}

	fmt.Printf("Fibonacci%d\n", fmax)
	mulprec.Display(&f1)

	fmt.Printf("Time: %s\n", time.Since(s))
}
