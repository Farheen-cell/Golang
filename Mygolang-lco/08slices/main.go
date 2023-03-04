package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice")
	var fruitList = []string{"Apple", "Tomato", "Mangos"}
	fmt.Printf("Type of fruitlist is : %T\n", fruitList)
	fruitList = append(fruitList, "banana", "graps")
	fmt.Println("fruit list is :", fruitList)
	//fruitList = append(fruitList[1:3])//last range will  always not inclusive
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 943
	highScore[2] = 787
	highScore[3] = 133
	//highScore[4] = 466 it give error like index out of bound
	highScore = append(highScore, 243, 546, 977)
	fmt.Println(highScore)

	fmt.Println("highScore new length", len(highScore))
	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slice based on index

	var courses = []string{"react", "js", "python", "java"}

	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) //for index element
	fmt.Println(courses)

}
