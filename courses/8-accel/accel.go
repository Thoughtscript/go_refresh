package main

import "fmt"

func GenDisplaceFn(accel float64, initVel float64, initDisplace float64) func (time float64) float64 {
	return func (time float64) float64 {
		a := (accel * (time*time))
		// Make sure to use float64 valid numbers here
		// Using a fraction and/or rational fraction here will actually calculate as 0!
		b := .5 * a
		c := initVel * time
		return b + c + initDisplace
	}
}

func main() {

	var accel float64
	var initVel float64
	var initDisplace float64

	fmt.Printf("Enter acceleration: ")
	fmt.Scan(&accel)
	fmt.Printf("Enter initial velocity: ")
	fmt.Scan(&initVel)
	fmt.Printf("Enter initial displacement: ")
	fmt.Scan(&initDisplace)

	var fn = GenDisplaceFn(accel, initVel, initDisplace)

	var timeA float64
	fmt.Printf("Enter first time: ")
	fmt.Scan(&timeA)

	fmt.Println("Calculating...")
	var resultA = fn(timeA)
	fmt.Printf("%f \n", resultA)

	var timeB float64
	fmt.Printf("Enter second time:")
	fmt.Scan(&timeB)

	fmt.Println("Calculating...")
	var resultB = fn(timeB)
	fmt.Printf("%f \n", resultB)
}
