package main

import (
	"fmt"
)

func decreaser() func() int {
	decreaseFrom := 10
	return func() int {
		decreaseFrom--
		return decreaseFrom
	}
}

func main() {
	dec := decreaser()
	fmt.Println(dec())
	//will return 8 because the variable decreaseFrom is bound to dec
	fmt.Println(dec())
	fmt.Println(dec())
}
