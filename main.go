package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

func main() {
	component := hello("John")

	/* SERVERING A COMPONENT */
	http.Handle("/", templ.Handler(component))

	http.Handle("/header", templ.Handler(headerTemplate("Tim")))

	/* RENDERING THE BUTTON ELEMENT */
	button("Click me").Render(context.Background(), os.Stdout)
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
