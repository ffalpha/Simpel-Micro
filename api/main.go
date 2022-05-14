package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Super secrent Information")
}

func isAuthrozied(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {

		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthrozied(homePage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Printf("Server")
	handleRequests()
}
