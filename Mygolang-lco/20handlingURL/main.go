package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=gfj35hjgjh"

func main() {
	fmt.Println("welcome to handling url")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)
	qparams := result.Query()
	fmt.Printf("The type of query param are : %T\n", qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])
	for _, val := range qparams {
		fmt.Println("Param is: ", val)

	}
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
