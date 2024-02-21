package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Item struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
	Date   string  `json:"date"`
}

var items = []Item{
	{Name: "Item 1", Amount: 10.34, Date: time.Now().Format("2006-01-02")},
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]

	var newItem Item

	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newItem.Date = date
	items = append(items, newItem)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func getItemsHandler(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]

	var result []Item
	for _, item := range items {
		if item.Date == date {
			result = append(result, item)
		}
	}

	json.NewEncoder(w).Encode(result)
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	date := mux.Vars(r)["date"]

	for i, item := range items {
		if item.Name == name && item.Date == date {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/items/{date}", createItemHandler).Methods("POST")
	r.HandleFunc("/items/{date}", getItemsHandler).Methods("GET")
	r.HandleFunc("/items/{date}/{name}", deleteItemHandler).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
