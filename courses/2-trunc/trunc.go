package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main () {

	fmt.Println("Enter floating point num: ")
	// Use bufio scanner to read input
	bscan := bufio.NewScanner(os.Stdin)

	// Continually reads input - try multiple times
	for bscan.Scan() {
		numStr := bscan.Text()
		if numStr != "" {

			// String to float conversion
			fl, pErr := strconv.ParseFloat(numStr, 8)
			// Float to int conversion
			trunc := int(fl)
			fmt.Println(trunc)

			// Exception handling
			if pErr != nil {
				fmt.Println("Parsing exception encountered: " + pErr.Error())
			}

			if bErr := bscan.Err(); bErr != nil {
				fmt.Println("Scanning exception encountered: " + bErr.Error())
			}
		}
	}
}
