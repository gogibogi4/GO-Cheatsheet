package main

import "fmt"

type Language struct {
	Name    string
	Version string
}

func Reversion(language *Language, callback chan *Language) {
	language.Version = "3.7"
	callback <- language
}

func main() {
	lang := new(Language)
	lang.Name = "Go"
	lang.Version = "1.10"
	fmt.Println(lang)
	ch := make(chan *Language)
	go Reversion(lang, ch)
	newVersion := <-ch
	fmt.Println(newVersion)
}
