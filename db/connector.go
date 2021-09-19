package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/joho/godotenv"
)

var db *sql.DB

// get from .env file
var server = goDotEnvVariable("SQLSERVER")
var port = goDotEnvVariable("SQLPORT")
var user = goDotEnvVariable("SQLUSER")
var password = goDotEnvVariable("SQLPASSWORD")
var database = goDotEnvVariable("SQLDATABASE")

func ConnectToDb() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
}

func ReadBauTable() (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT BauRef FROM dbo.BauTable;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var BauRef string

		// Get values from row.
		err := rows.Scan(&BauRef)
		if err != nil {
			return -1, err
		}

		fmt.Printf("BauRef: %s\n", BauRef)
		count++
	}

	return count, nil
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
