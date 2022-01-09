package auth

import (
	"bill-ly/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/store"
)

var authenticator auth.Authenticator
var cache store.Cache

func CreateToken(w http.ResponseWriter, r *http.Request){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "Madrix",
		"passcode" : "test",	
	})
	jwtToken, _:= token.SignedString([]byte("secret"))
    w.Write([]byte(jwtToken))
}

func ValidateUser(w http.ResponseWriter, r *http.Request, userInput models.UserLogin) {

}