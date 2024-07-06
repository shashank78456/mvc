package models

import (
	"database/sql"
	"context"
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func dsn() string {
	err := godotenv.Load()
    if (err != nil) {
        fmt.Println("Error loading .env file", err)
    }

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOST")
	database   := os.Getenv("DB_NAME")


    return fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, hostname, database)
}


func Connection() (*sql.DB, error) {  
	db, err := sql.Open("mysql", dsn())  
	if (err != nil) {  
		log.Printf("Error: %s when opening DB", err)
		return nil, err
	}

    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(20)
    db.SetConnMaxLifetime(time.Minute * 5)

    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    err = db.PingContext(ctx)
    if err != nil {
        log.Printf("Errors %s pinging DB", err)
        return nil, err
    }

	return db, err
}