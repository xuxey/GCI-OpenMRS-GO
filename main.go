package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Message object
type msg struct {
	Message string
}

// / endpoint
func homeLink(w http.ResponseWriter, r *http.Request) {
	user := &msg{Message: "Hello, World!"}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(b))
}

// /greet/{name} endpoint
func nameLink(w http.ResponseWriter, r *http.Request) {
	name := "Hello " + mux.Vars(r)["name"]
	user := &msg{Message: name}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(b))
}

func main() {
	// init router
	router := mux.NewRouter().StrictSlash(true)
	// register endpoints
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/greet/{name}", nameLink).Methods("GET")
	// print server status
	fmt.Println(http.ListenAndServe(":8088", router))
}
