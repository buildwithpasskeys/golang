package main

import (
	"fmt"
	"net/http"
)

func serve(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "hello\n")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", serve)
	fmt.Println("Serving on port: 8090. Press Ctrl-C to stop.")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
}
