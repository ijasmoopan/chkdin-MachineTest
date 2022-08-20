package respository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Can't acces env file")
	}
	driver := os.Getenv("POSTGRES_DRIVER")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DBNAME")

	dbString := fmt.Sprint(driver, "://", user, ":", password, "@", host, ":", port, "/", dbName, "?sslmode=disable")

	db, err := sql.Open(driver, dbString)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}
	return db
}