package main

import "fmt"

func main() {
	fmt.Println("Welcome to Type with Switch")
	PrintType(2)
	PrintType("farheen")
}

func PrintType(t interface{}) {
	switch v := t.(type) {
	case string:
		fmt.Println("Type: String")
	case int:
		fmt.Println("Type: int")
	default:
		fmt.Printf("Unkown Type %T", v)
	}
}
