package auth

import (
	"bill-ly/models"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/auth/strategies/bearer"
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

func ValidateUser(ctx context.Context, r *http.Request, userInput models.UserLogin) (auth.Info, error) {
	return nil, nil
}

func ValidateToken(ctx context.Context, r *http.Request, tokenString string) (auth.Info, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := auth.NewDefaultUser(claims["username"].(string), "", nil, nil)
		return user, nil
	}

	return nil, errors.New("invaled token")

}

func SetupGoGurardian() {
	authenticator = auth.New()
	cache = store.NewFIFO(context.Background(), time.Minute*10)

	// basicStrategy := basic.New(ValidateUser, cache)
	tokenStrategy := bearer.New(ValidateToken, cache)

	// authenticator.EnableStrategy(basic.StrategyKey, basicStrategy)
	authenticator.EnableStrategy(bearer.CachedStrategyKey, tokenStrategy)
}

func Middleware(next http.Handler) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Executing Auth Middleware")
        user, err := authenticator.Authenticate(r)
        if err != nil {
            code := http.StatusUnauthorized
            http.Error(w, http.StatusText(code), code)
            return
        }
		fmt.Println(user)
        fmt.Printf("User %s Authenticated\n", user.UserName())
        next.ServeHTTP(w, r)
    })
}