package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/joho/godotenv"
)

// https://docs.microsoft.com/en-us/azure/azure-sql/database/connect-query-go

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

// examples

// // postAlbums adds an album from JSON received in the request body.
// func PostAlbums(c *gin.Context) {
// 	var newAlbum models.Album

// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// // getAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func GetAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop over the list of albums, looking for
// 	// an album whose ID value matches the parameter.
// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

func GetData(table string) (string, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return "-1", err
	}

	tsql := fmt.Sprintf("SELECT * FROM %s;", table)

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return "-1", err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return "-1", err
	}

	result, err := convertToJson(rows, columns)

	return result, err
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

// https://www.sohamkamani.com/golang/json/
func convertToJson(rows *sql.Rows, columns []string) (string, error) {
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
