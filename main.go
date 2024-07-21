package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Define the Item struct
type Item struct {
	ID    int
	Name  string
	Price float64
}

var items []Item

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/items", itemsHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getItems(w, r)
	case http.MethodPost:
		createItem(w, r)
	case http.MethodPut:
		updateItem(w, r)
	case http.MethodDelete:
		deleteItem(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getItems(w http.ResponseWriter, r *http.Request) {
	// Convert items slice to JSON
	jsonItems, err := json.Marshal(items)
	if err != nil {
		http.Error(w, "Failed to marshal items to JSON", http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonItems)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	// Decode JSON request body into new item
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	// Assign ID to the new item
	newItem.ID = len(items) + 1

	// Add the new item to the items slice
	items = append(items, newItem)

	// Write success response
	w.WriteHeader(http.StatusCreated)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// Decode JSON request body into updated item
	var updatedItem Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	// Find the index of the item to be updated
	var found bool
	for i, item := range items {
		if item.ID == updatedItem.ID {
			items[i] = updatedItem
			found = true
			break
		}
	}

	// If the item is not found, return 404 Not Found
	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusOK)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Decode JSON request body into item ID to be deleted
	var itemID struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&itemID)
	if err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	// Find the index of the item to be deleted
	var found bool
	for i, item := range items {
		if item.ID == itemID.ID {
			// Remove the item from the slice
			items = append(items[:i], items[i+1:]...)
			found = true
			break
		}
	}

	// If the item is not found, return 404 Not Found
	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusOK)
}
