package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/joho/godotenv"
)

// https://docs.microsoft.com/en-us/azure/azure-sql/database/connect-query-go

// revisit this later for sql refinement
// https://www.calhoun.io/inserting-records-into-a-postgresql-database-with-gos-database-sql-package/

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

func PostNewItem(table string, requestBody []byte) ([]byte, error) {

	// decode JSON into an iterable structure
	c := make(map[string]json.RawMessage)
	e := json.Unmarshal(requestBody, &c)
	if e != nil {
		panic(e)
	}

	// create tsql string
	tsqlstart := fmt.Sprintf("INSERT INTO %s ( ", table)
	tsqlend := fmt.Sprintf("VALUES ( ")

	fmt.Println(len(c))

	i := 1

	for k, v := range c {
		if i != len(c) {
			tsqlstart = tsqlstart + string(k) + ", "
			tsqlend = tsqlend + strings.Replace(string(v), "\"", "'", -1) + ", "
		} else {
			tsqlstart = tsqlstart + string(k)
			tsqlend = tsqlend + strings.Replace(string(v), "\"", "'", -1)
		}
		i++
	}

	tsqlstart = tsqlstart + " ) "
	tsqlend = tsqlend + " )"

	tsql := tsqlstart + tsqlend

	// output result to STDOUT
	fmt.Println(tsql)

	_, err := db.Exec(tsql)
	if err != nil {
		panic(err)
	}

	return requestBody, err

}

func GetDataById(table string, id string) (string, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return "-1", err
	}

	tsql := fmt.Sprintf("SELECT * FROM %s WHERE Id = %s", table, id)

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
