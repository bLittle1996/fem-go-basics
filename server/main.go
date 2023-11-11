package main

import (
	"encoding/json"
	"fmt"
	"front-end-masters/go-basics/server/api"
	"html/template"
	"net/http"
	"os"
)

func main() {
	server := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))

	server.Handle("/", fs)

	server.Handle("/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type data struct {
			Name        string
			Description string
			Rarity      string
		}

		var templateData []data

		testData, err := os.ReadFile("./data/test.json")

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		err = json.Unmarshal(testData, &templateData)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		testTemplate, err := template.ParseFiles("./templates/index.tmpl.html")

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		testTemplate.Execute(w, templateData)
	}))

	server.Handle("/api/items", http.HandlerFunc(api.GetItems))

	err := http.ListenAndServe(":42069", server)

	if err != nil {
		fmt.Println(err)
	}
}
