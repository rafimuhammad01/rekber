package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB(dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode string) *sqlx.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to postgreSQL: %v", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping to postgreSQL: %v", err.Error())
	}

	return db
}
