package main

import (

	"fmt"
	"log"
	"net/http"
	//"strings"
	
	
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("Server started on: http://localhost:8000")
	//err := http.ListenAndServe(":8000", context.ClearHandler(http.DefaultServeMux)) // context to prevent memory leak
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("500 Internal server error", http.StatusInternalServerError) // internal server error
		return
	}
}