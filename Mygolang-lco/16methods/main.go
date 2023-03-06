package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")
	//no inharitance in golang no super no parants
	farheen := User{"farheen", "f@123", true, 29}
	fmt.Println(farheen)
	fmt.Printf("farheen details are: %+v\n", farheen)
	fmt.Printf("Name is %v and email is %v.", farheen.Name, farheen.Email)
	farheen.GetStatust()
	farheen.newMail() //only copy not actualy change for change in actuale we use pointer
	fmt.Printf("Name is %v and email is %v.", farheen.Name, farheen.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatust() {
	fmt.Println("Is user online", u.Status)

}
func (u User) newMail() {
	u.Email = "Farheen.ali108@gmail.com"
	fmt.Println("new mail", u.Email)
}
