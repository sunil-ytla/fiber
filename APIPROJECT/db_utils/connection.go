package db_utils

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq" // PostgreSQL driver
    "github.com/joho/godotenv"
)

var DB *sql.DB

func Init() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: No .env file found, using system environment variables")
    }

    // Read variables
    user := os.Getenv("PG_USER")
    password := os.Getenv("PG_PASSWORD")
    dbname := os.Getenv("PG_DBNAME")
    host := os.Getenv("PG_HOST")
    port := os.Getenv("PG_PORT")
    sslmode := os.Getenv("PG_SSLMODE")

    // Build DSN
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        host, port, user, password, dbname, sslmode,
    )

    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Failed to open DB:", err)
    }

    DB.SetMaxOpenConns(25)
    DB.SetMaxIdleConns(25)
    DB.SetConnMaxLifetime(0)

    if err = DB.Ping(); err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    log.Println("Database connection established")
}
