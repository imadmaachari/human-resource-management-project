package infrastructure

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

//Database struct
type Database struct {
	DB *sqlx.DB
}

//NewDatabase : Initialize and return pg db
func NewDatabase() Database {
	LoadEnv()
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASS, DBNAME)
	fmt.Println(URL)
	db, err := sqlx.Open("postgres", URL)
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connection established .")
	return Database{
		DB: db,
	}
}
