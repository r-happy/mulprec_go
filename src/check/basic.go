package check

import (
	"fmt"

	"github.com/r-happy/mulprec_go/src/mulprec"
)

func BasicCheck() {
	fmt.Println("Basic Check")
	Add()
	Sub()
	Increment()
	Decrement()
	ShiftLeft()
	ShiftRight()
	Multiple()
	Divide()
	Inverse()
}

func Add() {
	fmt.Println("Add")

	var a mulprec.NUMBER
	var b mulprec.NUMBER
	var c mulprec.NUMBER

	mulprec.SetInt(&a, 1)
	mulprec.SetInt(&b, 2)
	mulprec.Clear(&c)

	fmt.Printf("a: ")
	mulprec.Display(&a)
	fmt.Println("")
	fmt.Printf("b: ")
	mulprec.Display(&b)
	fmt.Println("")

	mulprec.Add(&a, &b, &c)

	fmt.Printf("c: ")
	mulprec.Display(&c)
	fmt.Println("")
}

func Sub() {
	fmt.Println("Sub")

	var a mulprec.NUMBER
	var b mulprec.NUMBER
	var c mulprec.NUMBER

	mulprec.SetInt(&a, 3)
	mulprec.SetInt(&b, 1)
	mulprec.Clear(&c)

	fmt.Printf("a: ")
	mulprec.Display(&a)
	fmt.Println("")
	fmt.Printf("b: ")
	mulprec.Display(&b)
	fmt.Println("")

	mulprec.Sub(&a, &b, &c)

	fmt.Printf("c: ")
	mulprec.Display(&c)
	fmt.Println("")
}

func Increment() {
	fmt.Println("Increment")

	var a mulprec.NUMBER
	var b mulprec.NUMBER

	mulprec.SetInt(&a, 1)
	mulprec.Clear(&b)

	fmt.Printf("a: ")
	mulprec.Display(&a)
	fmt.Println("")

	mulprec.Increment(&a, &b)

	fmt.Printf("b: ")
	mulprec.Display(&b)
	fmt.Println("")
}

func Decrement() {
	fmt.Println("Decrement")

	var a mulprec.NUMBER
	var b mulprec.NUMBER

	mulprec.SetInt(&a, 6)
	mulprec.Clear(&b)

	fmt.Printf("a: ")
	mulprec.Display(&a)
	fmt.Println("")

	mulprec.Decrement(&a, &b)

	fmt.Printf("b: ")
	mulprec.Display(&b)
	fmt.Println("")
}

func ShiftLeft() {
	fmt.Println("ShiftLeft")

	var a mulprec.NUMBER
	var b mulprec.NUMBER

	mulprec.SetInt(&a, 1)
	mulprec.Clear(&b)

	fmt.Printf("a: ")
	mulprec.Display(&a)
	fmt.Println("")

	mulprec.ShiftLeft(&a, &b, 2)

	fmt.Printf("b: ")
	mulprec.Display(&b)
	fmt.Println("")
}

func ShiftRight() {
	fmt.Println("ShiftRight")

	var a mulprec.NUMBER
	var b mulprec.NUMBER

	mulprec.SetInt(&a, 1)
	mulprec.Clear(&b)

	fmt.Printf("a: ")
	mulprec.Display(&a)
	fmt.Println("")

	mulprec.ShiftLeft(&a, &b, 5)
	mulprec.ShiftRight(&b, &a, 3)

	fmt.Printf("b: ")
	mulprec.Display(&a)
	fmt.Println("")
}

func Multiple() {
	fmt.Println("Multiple")

	var a mulprec.NUMBER
	var b mulprec.NUMBER
	var c mulprec.NUMBER

	mulprec.SetInt(&a, 123456789123456789)
	mulprec.SetInt(&b, 123456789123456789)
	mulprec.Clear(&c)

	fmt.Printf("a: ")
	mulprec.Display(&a)
	fmt.Println("")
	fmt.Printf("b: ")
	mulprec.Display(&b)
	fmt.Println("")

	mulprec.Multiple(&a, &b, &c)

	fmt.Printf("c: ")
	mulprec.Display(&c)
	fmt.Println("")

}

func Divide() {
	fmt.Println("Divide")

	var a mulprec.NUMBER
	var b mulprec.NUMBER
	var c mulprec.NUMBER

	mulprec.SetInt(&a, 121932631112635269)
	mulprec.SetInt(&b, 123456789)
	mulprec.Clear(&c)

	fmt.Printf("a: ")
	mulprec.Display(&a)
	fmt.Println("")
	fmt.Printf("b: ")
	mulprec.Display(&b)
	fmt.Println("")

	mulprec.Divide(&a, &b, &c)

	fmt.Printf("c: ")
	mulprec.Display(&c)
	fmt.Println("")
}

func Inverse() {
	fmt.Println("Inverse")

	var a mulprec.NUMBER
	var b mulprec.NUMBER

	mulprec.SetInt(&a, 123456789)
	mulprec.Clear(&b)

	fmt.Printf("a: ")
	mulprec.Display(&a)
	fmt.Println("")

	mulprec.Inverse(&a, &b, 100)

	fmt.Printf("b: ")
	mulprec.Display(&b)
	fmt.Println("")
}
