package main

import "fmt"

func main() {
	var fname string
	fmt.Println("Hello world")
	fmt.Scanln(&fname)
	fmt.Println("My name is " + fname)

}
