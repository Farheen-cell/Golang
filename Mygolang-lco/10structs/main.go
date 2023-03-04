package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")
	//no inharitance in golang no super no parants
	farheen := User{"farheen", "f@123", true, 29}
	fmt.Println(farheen)
	fmt.Printf("farheen details are: %+v\n", farheen)
	fmt.Printf("Name is %v and email is %v.", farheen.Name, farheen.Email)

}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}
