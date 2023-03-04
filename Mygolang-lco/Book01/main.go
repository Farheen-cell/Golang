package main

import (
	"fmt"
	"net/http"
)

//	func handleFunc(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprint(w, "<h1>Welcome to My awesome site!</h1")
//	}
func pathHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, r.URL.Path)
	if r.URL.Path == "/" {
		homeHandler(w, r)
	} else if r.URL.Path == "/contact" {
		constantHandler(w, r)
	} else {
		http.Error(w, "Page not found", http.StatusNotFound)

	}
}

func constantHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jon@calh")
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}
func main() {
	//http.HandleFunc("/", handleFunc)
	http.HandleFunc("/", pathHandler)

	fmt.Println("Starting the server on 3000...")
	http.ListenAndServe(":3000", nil)
}
