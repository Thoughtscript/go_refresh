package main

import (
	"encoding/json"
	"fmt"
)

type PersonalData struct{
	name string
	address string
}

func main()  {
	name := ""
	address := ""

	// Single entry - reads into the reference
	fmt.Println("Enter an address!")
	fmt.Scan(&address)
	fmt.Println("Enter a name!")
	fmt.Scan(&name)

	// Make a map not a struct
	m := make(map[string]string)
	m["name"] = name
	m["address"] = address
	fmt.Println("You entered - name: " + m["name"] + " address: " + m["address"])

	// Marshals from map and prints
	jm, mErr := json.Marshal(m)
	fmt.Printf("Marshalled from map: %s \n", jm)

	// Exception handling
	if mErr != nil {
		fmt.Println("Exception encountered " + mErr.Error())
	}

}