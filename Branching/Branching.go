package main

import "fmt"

func main() {
	fmt.Println(ifDemo(5, 10))

	fmt.Println(switchDemo("Bojan"))
}

func ifDemo(x int, y int) (result string) {
	if x > y {
		result = "X is greater than Y"
	} else if x < y {
		result = "Y is greater than X"
	} else {
		result = "They are the same"
	}

	return
}

func switchDemo(name string) (greeting string) {
	switch name {
	case "Igor":
		greeting = "Hello Igor"
	case "Eve":
		greeting = "Heloo miss"
	case "Bobi", "Mirko":
		greeting = "Hello dude"
	default:
		greeting = "Hello " + name
	}

	return greeting
}
