package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type user struct {
	Name  string `json:"full_name"`
	Email string `json:"email_address"`
}

type apiHandler struct{}

func (apiHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	u := user{Name: "Imad Elmahrad", Email: "imad.elmahrad98@gmail.com"}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(u)
}

func homeHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello from home handler %s", req.URL)
}

func logger(fn http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		fn(rw, req)
		end := time.Since(start)
		fmt.Printf("%s %s processing time %s\n", req.Method, req.URL, end)
	}
}

func main() {
	mux := http.DefaultServeMux

	mux.Handle("/api", apiHandler{})
	mux.HandleFunc("/home", logger(homeHandler))

	http.ListenAndServe(":3000", mux)
}
