package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")
	parsentTime := time.Now()

	fmt.Println(parsentTime)
	fmt.Println(parsentTime.Format("01-02-2006 15:04:05 Monday"))
	cratedDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(cratedDate)
	fmt.Println(cratedDate.Format("01-02-2016 15:04:05 Monday"))
}
