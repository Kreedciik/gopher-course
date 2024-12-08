package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Language", "en")
		fmt.Fprintf(w, "Hello Gopher")
	})

	http.ListenAndServe(":6000", nil)
}
