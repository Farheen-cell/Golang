package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Map")
	laguages := make(map[string]string)
	laguages["js"] = "javascript"
	laguages["rb"] = "ruby"
	laguages["py"] = "python"

	fmt.Println(laguages) //come in sorted form
	fmt.Println("js short for :", laguages["js"])

	delete(laguages, "rb")
	fmt.Println("List of all laguages ", laguages)

	for key, value := range laguages {
		fmt.Printf("for key %v ,value id %v\n", key, value)
	}
	//convarte map in json
	a := make(map[int]string)
	a[1] = "John"
	j, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
	} else {
		fmt.Println(string(j))
	}
	//convart json in map

	z := `{"1":"John"}`
	var b map[string]interface{}
	json.Unmarshal([]byte(z), &b)
	fmt.Println(b)
	//employee json in map

	e := `{"1":{"Name":"John"},
	"2":{"Name":"Farheen"}
	}`
	var m map[int]employee
	json.Unmarshal([]byte(e), &m)
	fmt.Println(m)

}

type employee struct {
	Name string
}
