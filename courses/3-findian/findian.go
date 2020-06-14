package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {

	fmt.Println("Enter string: ")
	// Use bufio scanner to read input
	bscan := bufio.NewScanner(os.Stdin)

	// Continually reads input - try multiple times
	for bscan.Scan() {
		line := bscan.Text()
		if line != "" {
			hasA := false

			// Loop and iterate through string
			for i := 0; i < len(line); i++ {
				if string(line[i]) == "a" || string(line[i]) == "A" {
					hasA = true
				}
			}

			// Boolean checks to see start and end chars
			iStart := string(line[0]) == "i" || string(line[0]) == "I"
			nEnd := string(line[len(line)-1]) == "n" || string(line[len(line)-1]) == "N"

			// Go has no ternary operator
			if hasA && nEnd && iStart { fmt.Println("Found!")
			} else { fmt.Println("Not Found!") }

			// Exception handling
			if bErr := bscan.Err(); bErr != nil {
				fmt.Println("Scanning exception encountered: " + bErr.Error())
			}
		}
	}
}
