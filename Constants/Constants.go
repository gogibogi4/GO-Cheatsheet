package main

import "fmt"

const (
	Language   = "English"
	Type       = "Spoken"
	Popularity = "High"
)

const (
	A = 1
	B = 2
	C = 3
	D = iota
	E = iota
	F
	G = iota
)

func main() {
	fmt.Println(Language, Type, Popularity)
	fmt.Printf("A: %d, B: %d, C: %d, D: %d, E: %d, F: %d, G: %d", A, B, C, D, E, F, G)
}
