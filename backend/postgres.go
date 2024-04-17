package backend

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var PGBackend *PostgresBackend

type PostgresBackend struct {
	db *sql.DB
}

func Init() {
	// connect to database
	connectionString := "host=localhost port=5432 user=root password=123456 dbname=test sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database successfully")

	// initialize tables
	file, _ := os.ReadFile("init.sql")
	_, err = db.Exec(string(file))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Initialzed tables successfully")

	PGBackend = &PostgresBackend{db: db}
}

func Close() {
	PGBackend.db.Close()
}
