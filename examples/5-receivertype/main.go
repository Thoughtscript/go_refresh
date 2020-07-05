package main

// Part of same package
import (
	data "./receiver"
	"fmt"
)

// GoLang generic
// The empty interface specifies that any type may be passed
func golangGeneric(val interface{}) {
	fmt.Println(val)
}

func main() {
	a := data.Point{1, 2, 3}
	fmt.Println(a)

	b := data.NewPoint(3,4,5)
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(*b)

	c := data.AddCoordinates(a)
	fmt.Println(c)

	d := a.AdditionReceiver()
	fmt.Println(d)

	golangGeneric(a)
}