package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	// Use MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// Storage is a main object that provides database functionality.
type Storage struct {
	db *sql.DB
}

// DatabaseConfig contains necessary configuration strings used for database initialization.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

const (
	// RetryCooldown is amount of second the system will wait before next database ping.
	RetryCooldown = 1

	// ConnectionTimeout total amount of time after which system will terminate with panic.
	ConnectionTimeout = 20
)

// ConnectDatabase tries to connect to the database and returns Storage, otherwise panic after a timeout.
func ConnectDatabase(dbConfig DatabaseConfig) *Storage {
	connString := constructConnectionString(dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Port, dbConfig.Name)
	db := initDatabase(connString)
	return &Storage{db: db}
}

func constructConnectionString(host string, user string, password string, port string, db string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
}

func initDatabase(conn string) *sql.DB {
	var db *sql.DB
	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeout*time.Second)
	defer cancel()
	var err error
	db, err = sql.Open("mysql", conn)
	if err != nil {
		panic("server: storage: could not open database connection")
	}
	for {
		select {
		case <-ctx.Done():
			panic(fmt.Sprintf("server: storage: couldn't connect to database in %d seconds.", ConnectionTimeout))
		default:
			err = pingDatabase(db)
			if err == nil {
				log.Printf("server: connected to database.")
				return db
			}
			log.Printf("server: storage: cannot connect to database, retrying in %d seconds.", RetryCooldown)
			time.Sleep(RetryCooldown * time.Second)
		}
	}
}

func pingDatabase(db *sql.DB) error {
	err := db.Ping()
	return err
}
