package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"reflect"
	"strings"
)

// Order the data structure for our eCommerce orders
type Order struct {
	Name   string
	Orders [3]int
}

func main() {
	fmt.Println("Hello World")

	possiblyMakeDirectory("orders")

	http.HandleFunc("/api/", apiHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func possiblyMakeDirectory(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		returnJSONResponse(w, r)
	case http.MethodPost:
		handleJSONRequest(w, r)
	}
}

func returnJSONResponse(w http.ResponseWriter, r *http.Request) {
	var order Order

	fmt.Println(r.Method)

	body, _ := ioutil.ReadFile("testdata.json")

	json.Unmarshal(body, &order)

	returnJSON(w, order)
}

func handleJSONRequest(w http.ResponseWriter, r *http.Request) {
	var order Order

	json.Unmarshal([]byte(r.FormValue("data")), &order)

	title := fmt.Sprintf("%s.json", order.Name)

	title = strings.Join(strings.Split(title, " "), "_")

	body, _ := json.Marshal(order)

	err := ioutil.WriteFile(fmt.Sprintf("orders/%s", title), body, 0777)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	returnJSON(w, order)
}

func returnJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}
