package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("The value is %v and index is %v\n", day, index)
	}

	rougueValu := 1

	for rougueValu < 10 {
		// if rougueValu == 5 {
		// 	break
		// }
		if rougueValu == 2 {
			goto loc

		}
		if rougueValu == 5 {
			rougueValu++
			continue

		}
		fmt.Println("Value is : ", rougueValu)
		rougueValu++
	}

loc:
	fmt.Println("Jumping to goto eaxample")
}
