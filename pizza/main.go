package main

import (
	"bill-ly/routes/auth"
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

func test(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	x := req.Form
	resp := `{"message" : ` + `"` + x["test"][0] + `"}`
	w.Write([]byte(resp))
}

func main() {
	mux := mux.NewRouter() 
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/post", test)
	mux.HandleFunc("/auth/token", auth.CreateToken).Methods("GET")
	fmt.Println("server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
