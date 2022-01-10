package main

import (
	"bill-ly/routes"
	"bill-ly/routes/auth"
	"bill-ly/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var mySigningKey = []byte("shlokisawesome")

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte(`{"message" : "hello"}`))
}

func main() {
	auth.SetupGoGurardian() // Setup go guardian
	DB := utils.InitDB()
	h := routes.New(DB)
	mux := mux.NewRouter() 
	mux.HandleFunc("/", auth.Middleware(http.HandlerFunc(handler))).Methods("GET")
	mux.HandleFunc("/auth/register", h.Register).Methods("POST")
	mux.HandleFunc("/auth/token", auth.CreateToken).Methods("GET")
	fmt.Println("server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
