package user

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func GeneratingToken(id int) (string, error) {

	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Println("Couldn't access env file")
	}
	key := os.Getenv("SECRET_KEY")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Println("Error while generating token")
		return "", err
	}
	return tokenString, nil
}