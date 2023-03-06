package main

import "fmt"

func main() {

	defer fmt.Println("Hello")
	defer fmt.Println("Hellotwo") //if bunch of defer it exicute in LIFO form
	fmt.Println("farheen")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
