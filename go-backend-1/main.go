package main

import (
	"fmt"
	"net/http"

	"github.com/AlexDiwong/build-testing/go-library-1/arithmetic"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!\n")
		fmt.Fprint(w, arithmetic.Add(1, 1))
	})

	http.ListenAndServe(":8080", nil)
}
