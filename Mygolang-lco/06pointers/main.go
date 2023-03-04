package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers ")

	//var ptr *int
	//fmt.Println("value of pointer is ", ptr)
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("value of atual pointer is ", ptr)
	fmt.Println("value of atual pointer is ", *ptr)
	*ptr = *ptr + 2
	fmt.Println(myNumber)

}
