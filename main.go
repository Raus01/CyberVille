package main

import (

	
	"log"
	"net/http"
	
	
	
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	
	http.HandleFunc("/", indexHandler)
	log.Println("Server started on: http://localhost:8000")
	//err := http.ListenAndServe(":8000", context.ClearHandler(http.DefaultServeMux)) // context to prevent memory leak
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("500 Internal server error", http.StatusInternalServerError) // internal server error
		return
	}
}