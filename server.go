package main
// Just a test comment

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Completed: \n")
	name := r.FormValue("name")
	password := r.FormValue("password")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Password = %s\n", password)
	fmt.Printf("Name = %s\n ", name)
	fmt.Printf("Password = %s\n ", password)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
