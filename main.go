package main

import (
	"fmt"
	"log"
	"sus-kpi-backend/controllers"
	"sus-kpi-backend/db"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectToDb()

	count, err := db.ReadBauTable()
	if err != nil {
		log.Fatal("Error reading bau benchmarks: ", err.Error())
	}
	fmt.Printf("Read %d row(s) successfully.\n", count)

	// Routes
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.POST("/albums", controllers.PostAlbums)

	router.Run("localhost:8080")
}
