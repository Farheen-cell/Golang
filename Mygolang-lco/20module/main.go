package main

import (
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome to go mod")
	greeder()
	//r := mux.NewRouter()
	//r.HandleFunc("/", serverHome).Methods("GET")
	//log.Fatal(http.ListenAndServe(":4000", r))
}
func greeder() {
	fmt.Println("Hey there mod users")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome tomy server</h1>"))

}
