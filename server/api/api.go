package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const secretDataFilePath = "./data/test.json"

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Rarity      string `json:"rarity"`
}

func loadDataFromFile() ([]Item, error) {
	var items []Item

	fileBytes, err := os.ReadFile(secretDataFilePath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileBytes, &items)

	if err != nil {
		return nil, err
	}

	return items, nil
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	items, err := loadDataFromFile()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"Internal Server Error\"}"))
		return
	}

	if r.URL.Query().Has("name") {
		getItem(w, r, items)
		return
	}

	err = json.NewEncoder(w).Encode(items)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"Internal Server Error\"}"))
	}
}

func PostItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	newItems, err := addItem(item)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = writeItemsToFile(newItems)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(item)
}

func addItem(item Item) ([]Item, error) {
	items, err := loadDataFromFile()

	if err != nil {
		return nil, err
	}

	items = append(items, item)

	return items, nil
}

func writeItemsToFile(items []Item) error {
	itemsJsonBytes, err := json.Marshal(items)

	if err != nil {
		return err
	}

	return os.WriteFile(secretDataFilePath, itemsJsonBytes, 0666)
}

func getItem(w http.ResponseWriter, r *http.Request, items []Item) {

	name := r.URL.Query().Get("name")
	var matchingItem *Item

	for _, item := range items {
		if item.Name == name {
			matchingItem = &item
			break
		}
	}

	if matchingItem == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"message\": \"Not Found\"}"))
		return
	}

	err := json.NewEncoder(w).Encode(matchingItem)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"Internal Server Error\"}"))
	}
}
