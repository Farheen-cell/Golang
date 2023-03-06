package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("welcome to web request")
	const url = "http://lco.dev"
	responce, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Responce is of type : %T\n", responce)
	defer responce.Body.Close()
	databyte, err := ioutil.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}
	contant := string(databyte)
	fmt.Println(contant)
}
