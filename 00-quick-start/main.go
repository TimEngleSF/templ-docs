package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	// Assign hello component with the string "John"
	// This will generate html code that contains the string "John"
	component := hello("John")

	// Server templ component on root
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
