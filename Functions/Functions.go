package main

import (
	"fmt"
)

func main() {
	simplest()
	fmt.Println(simpleReturn("This was the message"))

	ok, err := multipleReturn("test")

	fmt.Println(ok)
	fmt.Println(err)

	okay, error := variableNumberArgs("msg", "one", "two")

	fmt.Println(okay)
	fmt.Println(error)
}

func variableNumberArgs(message string, rest ...string) (ok int, err string) {
	fmt.Println("The message was" + message)
	fmt.Println("The rest were:", rest)

	ok = len(rest)
	err = "No error"

	return
}

func multipleReturn(message string) (ok string, err string) {
	ok = "Got message: " + message
	err = "No error"

	return ok, err
}

func simpleReturn(message string) string {
	return "Got message " + message
}

func simplest() {
	fmt.Println("Simplest")
}
