package main

import (
	"log"
	"net/http"
	
	
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func ad(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin.html")
}
func us(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "user.html")
}

func main() {

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	//20.91.189.150
	address := "20.91.189.150:8000"

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/admin", ad)
	http.HandleFunc("/user", us)
	log.Printf("Server started on: http://%s\n", address)
	//err := http.ListenAndServe(":8000", context.ClearHandler(http.DefaultServeMux)) // context to prevent memory leak
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
		return
	}
}