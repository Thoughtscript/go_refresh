package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	a := make([]int, 3, 3)
	idx := 0
	entries := 0
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
				if idx < 3 {
					a[idx] = numInt
					idx++
				} else {
					a = append(a, numInt)
				}
				entries++

				// Easier to copy and sort the copy
				b := make([]int, len(a), len(a))
				copy(b,a)
				sort.Ints(b)
				fmt.Printf("Array: %v Capacity: %d Length %d Entries %d \n", b, cap(b), len(b), entries)

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
