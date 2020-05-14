package main

import (
	"fmt"
)

func main() {
	//goish
	var myGoishSlice []int
	//initialize (type, initial_cap, max_cap)
	myGoishSlice = make([]int, 3, 5)
	myGoishSlice[0] = 1
	myGoishSlice[1] = 2
	myGoishSlice[2] = 3

	fmt.Println(myGoishSlice)

	//append
	myGoishSlice = append(myGoishSlice, 4)
	fmt.Println(myGoishSlice)

	//shorthand
	myShortSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(myShortSlice)

	//delete
	myShortSlice = append(myShortSlice[:1], myShortSlice[3:]...)
	fmt.Println(myShortSlice)
}
