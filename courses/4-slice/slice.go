package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := make([]int, 3, 3)
	idx := 0
	fmt.Println("Enter an integer or X to terminate the program: ")
	// Use bufio scanner to read input
	bscan := bufio.NewScanner(os.Stdin)
	// Continually reads input - try multiple times
	for bscan.Scan() {
		line := bscan.Text()
		if line != "" {
			if line == "X" {
				fmt.Println("Exiting program!")
				os.Exit(0)
			} else {
				pInt, iErr := strconv.ParseInt(line, 10, 64)
				numInt := int(pInt)
				if idx < len(a) {
					a[idx] = numInt
					idx++
				} else {
					a = append(a, numInt)
				}
				fmt.Printf("Array: %v Capacity: %d Length %d \n", a, cap(a), len(a))
				// Exception handling
				if iErr != nil {
					fmt.Println("Int formatting exception encountered: " + iErr.Error())
				}
				if bErr := bscan.Err(); bErr != nil {
					fmt.Println("Scanning exception encountered: " + bErr.Error())
				}
			}
		}
	}

}
