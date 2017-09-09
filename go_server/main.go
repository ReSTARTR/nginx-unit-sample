package main

import (
	"fmt"
	"net/http"

	// "github.com/nginx/unit/src/go/unit"
	"github.com/ReSTARTR/unit/src/go/unit"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go")
	})
	unit.ListenAndServe(":8080", nil)
}
