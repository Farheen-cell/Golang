package main

import "fmt"

func main() {
	fmt.Println("welcome to if else")
	logincount := 23

	var result string

	if logincount < 10 {
		result = "regular user"
	} else if logincount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count "
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Numbere is less then 10")
	} else {
		fmt.Println("Number is not less then 10")
	}

	//if err!=nil{

}
