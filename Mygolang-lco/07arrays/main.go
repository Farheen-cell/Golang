package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Mango"
	fmt.Println("Fruits list is :", fruitList)
	fmt.Println("length of list is :", len(fruitList))

	var vagList = [5]string{"patato", "beans", "masroooms"}
	fmt.Println("vag list is: ", vagList)
	fmt.Println("length of list", len(vagList))
}
