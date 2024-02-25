package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const date_format string = "02-01-2006"

type Item struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
	Date   string  `json:"date"`
}

var items = []Item{
	{Name: "Curd", Amount: 88.00, Date: time.Now().Format(date_format)},
	{Name: "Eggs", Amount: 54.00, Date: time.Now().Format(date_format)},
	{Name: "Peanut Butter", Amount: 125.00, Date: time.Now().Format(date_format)},
	{Name: "Butter", Amount: 58.00, Date: time.Now().Format(date_format)},
	{Name: "Lemon", Amount: 52.00, Date: time.Now().Format(date_format)},
	{Name: "Bread", Amount: 54.00, Date: time.Now().Format(date_format)},
	{Name: "Face Wash", Amount: 368.00, Date: time.Now().Format(date_format)},
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

	w.Header().Set("Content-Type", "application/json")
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

	w.Header().Set("Content-Type", "application/json")
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
