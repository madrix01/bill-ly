package routes

import (
	"bill-ly/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	uuid "github.com/google/uuid"
)

func (h Handler)Register(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	
	var newUser models.User
	newUser.Id = uuid.NewString()  
	json.Unmarshal(body, &newUser)
	fmt.Println(newUser)
	if result := h.DB.Create(&newUser); result.Error != nil {
        fmt.Println(result.Error)
    }

    // Send a 201 created response
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("New user created")
}