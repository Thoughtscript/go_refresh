package main

// Part of same package
import (
	data "./receiver"
	"fmt"
)

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
}