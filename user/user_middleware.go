package user

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ijasmoopan/chkdin-MachineTest/models"
	"github.com/joho/godotenv"
)

func (repo *Repo) IsUserAuthorized(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var message interface{}
		err := godotenv.Load("./config/.env")
		if err != nil {
			log.Println("Couldn't access env file")
			return
		}
		key := os.Getenv("SECRET_KEY")

		var reqToken string
		if r.Header["Authorization"] != nil {
			reqToken = r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]
		} else {
			message = "Token not found"
			responseBadRequest(w, message)
			return
		}
		token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Error in token")
				// return "Error in token", errors.New("Error")
			}
			return []byte(key), nil
		})
		log.Println("Token:", token)
		if err != nil {
			message = "Unauthorized token"
			responseBadRequest(w, message)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			id := claims["id"]
			user, err := repo.user.DBAuthUser(id.(float64))
			if err != nil {
				message = "Please Login"
				responseBadRequest(w, message)
				return
			}
			ctx := context.WithValue(r.Context(), models.CtxKey{}, user)
			handler.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
