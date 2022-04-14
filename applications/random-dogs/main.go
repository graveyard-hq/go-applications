package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Dog struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func dogHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var dog Dog
	err = json.Unmarshal(body, &dog)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, dog.Message)
}

func main() {
	http.HandleFunc("/", dogHandler)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
