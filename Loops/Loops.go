package main

import (
	"fmt"
)

type Laguages struct {
	name       string
	isCompiled bool
}

func main() {
	//for
	for i := 1; i < 10; i++ {
		fmt.Println("The i is ", i)
	}

	//while
	var j int = 1

	for j <= 10 {
		fmt.Println(j)
		j++
	}

	slice := []Laguages{
		{"Python", false},
		{"GO", true},
		{"C", true},
	}

	for n, c := range slice {
		fmt.Println("Here we go: ", n, c.name, c.isCompiled)
	}
}
