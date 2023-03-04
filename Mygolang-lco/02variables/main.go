package main

import "fmt"

//jwtToken :=3000 it will give error like outside main declaring variable by :
// is not aloud we need to inasalise variable with type

const LoginToken string = "sdfsafs" //if we use name of variable staring from capital latter so i will become public variable

func main() {
	//fmt.Println("variables")
	//string example
	var username string = "farheen"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	//boolean example
	var isloggedIn bool = true
	fmt.Println(isloggedIn)
	fmt.Printf("Variable is of type: %T \n", isloggedIn)

	//Integer example
	var smallval uint16 = 256
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	//flote example
	var smallfloat float64 = 255.3325345643635345
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	//defullt values and some aliases
	var anotherval int
	fmt.Println(anotherval)
	fmt.Printf("Variable is of type: %T \n", anotherval)

	var anotherstring string
	fmt.Println(anotherstring)
	fmt.Printf("Variable is of type: %T \n", anotherstring)

	//implicit type

	var webside = "farheen.omniful.com"
	fmt.Println(webside)

	//no var style

	numberOfuser := 30000
	fmt.Println(numberOfuser)

	numberOfuser2 := 30000.0
	fmt.Println(numberOfuser2)

	//use of public variable

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
