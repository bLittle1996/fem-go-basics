package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Rarity      string `json:"rarity"`
}

func loadDataFromFile() ([]Item, error) {
	var items []Item

	fileBytes, err := os.ReadFile("./data/test.json")

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
		getItem(w, r, &items)
		return
	}

	err = json.NewEncoder(w).Encode(items)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"Internal Server Error\"}"))
	}

}

func getItem(w http.ResponseWriter, r *http.Request, items *[]Item) {

	name := r.URL.Query().Get("name")
	var matchingItem *Item

	for _, item := range *items {
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
