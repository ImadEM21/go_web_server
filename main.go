package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Hello from go server [%s]", req.URL)
	})
	http.ListenAndServe(":3000", nil)
}
