package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file in golang")
	content := "This needs to go in a file -LearnCodeOnline.com"

	file, err := os.Create("./mylcogofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	chekNilError(err)

	length, err := io.WriteString(file, content)

	chekNilError(err)
	fmt.Println("length is :", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	chekNilError(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

func chekNilError(err error) {
	if err != nil {
		panic(err)
	}
}
