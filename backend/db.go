package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// connectDB retries until the PostgreSQL server is ready
func connectDB() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// default connection if env variable is not set
		dsn = "host=db user=postgres password=postgres dbname=paris sslmode=disable"
	}

	var db *sql.DB
	var err error

	// Retry loop: wait for DB to be ready
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			fmt.Println("Erreur ouverture DB:", err)
		} else {
			err = db.Ping() // test connection
			if err == nil {
				break // success
			}
			fmt.Println("DB pas prête, retry...")
		}
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("impossible de se connecter à la DB: %v", err)
	}

	// Create table if it doesn't exist
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS arbres_remarquables (
		id TEXT PRIMARY KEY,
		nom_usuel TEXT,
		nom_latin TEXT,
		genre TEXT,
		espece TEXT,
		arr TEXT,
		lat DOUBLE PRECISION,
		lon DOUBLE PRECISION,
		hauteur DOUBLE PRECISION,
		url_pdf TEXT,
		url_photo1 TEXT,
		resume TEXT
	)`)
	if err != nil {
		return nil, fmt.Errorf("erreur création table: %v", err)
	}

	fmt.Println("✅ Base de données prête !")
	return db, nil
}
