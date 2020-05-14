package main

import "fmt"

type Languages []Language

type Language struct {
	Name    string
	Version string
}

type Reversionable interface {
	Reversion(newVersion string)
}

func (Language *Language) Reversion(newVersion string) {
	Language.Version = newVersion
}

func ReversionTo(v Reversionable) {
	v.Reversion("2.7")
}

func main() {
	languages := Languages{
		{"Python", "3.7"},
		{"Go", "1.11"},
	}

	fmt.Println(languages)
	ReversionTo(&languages[0])
	fmt.Println(languages)
}
