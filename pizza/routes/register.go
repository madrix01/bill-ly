package routes

import (
	"bill-ly/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"strings"

	uuid "github.com/google/uuid"
)

func (h Handler) Register(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var newUser models.User
	newUser.Id = uuid.NewString()
	json.Unmarshal(body, &newUser)

	// Username exsist
	var exsistUser models.User
	userNameExist := h.DB.First(&exsistUser, "username = ?", newUser.Username)
	emailExist := h.DB.First(&exsistUser, "email = ?", newUser.Email)
	if userNameExist.Error == nil || emailExist == nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error" : "user already exist!"}`))
		return
	}

	if !checkMail(newUser.Email) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error" : "invalid email!"}`))
		return
	}

	if result := h.DB.Create(&newUser); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("New user created")
}

func checkMail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func checkUsername(username string) bool {
	username = strings.TrimSpace(username)
	if len(username) < 5 {
		return false
	}
	return false
}
