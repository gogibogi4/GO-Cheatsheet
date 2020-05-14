package main

import "fmt"

func main() {
	var myString string = "Welcome to go programming!"
	var myStringPointer *string

	myStringPointer = &myString

	fmt.Println(myString, *myStringPointer, myStringPointer)

	*myStringPointer = "Hello"

	fmt.Println(myString, *myStringPointer, myStringPointer)
}
