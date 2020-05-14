package main

import "fmt"

type Languages struct {
	Name    string
	Version string
}

// If the pointer is not provided, the variable will be passed as a copy
func (language *Languages) Reversion(newVersion string) {
	language.Version = newVersion
}

func main() {
	languages := Languages{Name: "Python", Version: "3.7"}
	fmt.Println(languages)

	languages.Reversion("2.7")
	fmt.Println(languages)
}
