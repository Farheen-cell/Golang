package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter reting for my pizza")

	//comma ok or error ok syntex
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("type of rating is %T", input)

}
