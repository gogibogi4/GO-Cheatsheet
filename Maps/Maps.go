package main

import (
	"fmt"
)

func main() {
	// map[key_type]value_tpye
	var namePrefix map[string]string
	//initialize
	namePrefix = make(map[string]string)
	namePrefix["Bojan"] = "Mr "
	namePrefix["Naomi"] = "Ms"
	namePrefix["Florian"] = "Mr"

	var name string = "Bojan"

	fmt.Println(namePrefix[name])

	// shorthand notation
	prefix := map[string]string{
		"Bojan":   "Mr",
		"Naomi":   "Miss",
		"Florian": "Mr",
	}

	name = "Naomi"

	fmt.Println(prefix[name])
}
