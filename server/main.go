package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there from go1 ")
}

func main() {
	// http.HandleFunc("/", handler)
	// var port = os.Getenv("PORT")
	// log.Printf("Running go lang server on port: %s", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
