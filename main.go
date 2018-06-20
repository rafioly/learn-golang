package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var supermarket map[string]int

func main() {
	supermarket = make(map[string]int)
	router := mux.NewRouter()
	router.HandleFunc("/", GetRoot).Methods("GET")
	router.HandleFunc("/add_item", AddItem).Methods("POST")
	router.HandleFunc("/list_all", ListAll).Methods("GET")
	router.HandleFunc("/", DeleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetRoot : Handler for GET / endpoint
func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You're on ROOT")
}

// AddItem : Handler for POST /add_item endpoint
func AddItem(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params.Get("name")
	itemPrice, _ := strconv.Atoi(params.Get("price"))

	supermarket[name] = itemPrice
	fmt.Fprintf(w, "Item named %s with price %d has been successfully added", name, itemPrice)
}

// ListAll : Handler for GET /list_all endpoint
func ListAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(supermarket)
}

// DeleteItem : Handler for DELETE / endpoint
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	delete(supermarket, name)

	fmt.Fprintf(w, "Item named %s has been successfully removed", name)
}
