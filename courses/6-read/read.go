package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct{
	firstname string
	lastname string
}

func main() {

	// array
	db := make([]Name, 0, 100)

	// Enter filenames
	fileName := ""
	fmt.Println("Enter filename!")
	fmt.Scan(&fileName)

	// Open file
	file, err := os.Open(fileName)
	// Scan contents of file
	bscan := bufio.NewScanner(file)

	// Iterates through
	for bscan.Scan() {
		line := bscan.Text()
		arr := strings.Split(line, " ")
		// Reads lines into struct
		n := Name{arr[0], arr[1]}
		// Adds struct to array
		db = append(db, n)
	}

	fmt.Printf("Array of stucts: %s \n", db)

	// Iterates through array printing data
	for i := 0; i < len(db); i++ {
		fmt.Println(" First name: " + db[i].firstname + " Last name: " + db[i].lastname)
	}

	// Error handling
	if err != nil {
		fmt.Println("Exception encountered " + err.Error())
	}

}