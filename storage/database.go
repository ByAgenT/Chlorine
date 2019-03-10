package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	// Use Postgres driver
	_ "github.com/lib/pq"
)

const (
	// RetryCooldown is amount of second the system will wait before next database ping.
	RetryCooldown = 1

	// ConnectionTimeout total amount of time after which system will terminate with panic.
	ConnectionTimeout = 20
)

// ID represents serial identification number of object in storage.
type ID int

// DBStorage is a main object that provides database functionality.
type DBStorage struct {
	db *sql.DB
}

// Query prepares and exececutes SQL query and return rows fetched from the database.
func (s DBStorage) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.db.Query(query, args)
}

// QueryRow prepare and execute SQL query and expects to return only one row.
func (s DBStorage) QueryRow(query string, args ...interface{}) *sql.Row {
	return s.db.QueryRow(query, args)
}

// Exec prepares and executes SQL query and return summarized result of SQL statement execution.
func (s DBStorage) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.db.Exec(query, args)
}

// DatabaseConfig contains necessary configuration strings used for database initialization.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

// ConnectDatabase tries to connect to the database and returns Storage, otherwise panic after a timeout.
func ConnectDatabase(dbConfig DatabaseConfig) *DBStorage {
	connString := constructConnectionString(dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Port, dbConfig.Name)
	db := initDatabase(connString)
	return &DBStorage{db: db}
}

func constructConnectionString(host string, user string, password string, port string, db string) string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, db, user, password)
}

func initDatabase(conn string) *sql.DB {
	var db *sql.DB
	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeout*time.Second)
	defer cancel()
	var err error
	db, err = sql.Open("postgres", conn)
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
			log.Printf("server: storage: cannot connect to database: %s. retrying in %d seconds.", err, RetryCooldown)
			time.Sleep(RetryCooldown * time.Second)
		}
	}
}

func pingDatabase(db *sql.DB) error {
	err := db.Ping()
	return err
}
