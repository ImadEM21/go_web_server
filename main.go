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

func withLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(rw, req)
		end := time.Since(start)
		fmt.Printf("%s %s processing time %s\n", req.Method, req.URL, end)
	})
}

func main() {
	mux := http.DefaultServeMux

	mux.Handle("/api", withLogger(apiHandler{}))
	mux.Handle("/home", withLogger(http.HandlerFunc(homeHandler)))

	http.ListenAndServe(":3000", mux)
}
