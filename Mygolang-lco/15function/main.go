package main

import "fmt"

func main() {
	fmt.Println("welcome to function")
	proResult, promsg := proAdder(2, 4, 6, 8)
	fmt.Println(proResult, promsg)

}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "pro result"
}
