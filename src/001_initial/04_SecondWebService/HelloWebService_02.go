package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	_ "github.com/gorilla/mux"
)
type API struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request) {
	// _urlParams := mux.Vars(r)
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request){
		
		
		
		
		message := API{"Hello, world"}
		output , err := json.Marshal(message)

		if(err != nil) {
			fmt.Println("Something went wrong")
		}
		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":8080", nil)
}