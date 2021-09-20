package main

import (
	"sus-kpi-backend/controllers"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/gin-gonic/gin"
)

func main() {

	// Routes
	router := gin.Default()
	router.GET("/bauBenchmarks", controllers.GetBauBenchmarks)
	router.GET("/projects", controllers.GetProjects)
	router.GET("/bauBenchmarks/:id", controllers.GetBauBenchmarkById)
	router.GET("/projects/:id", controllers.GetProjectById)
	// router.POST("/projects", controllers.CreateProject)

	router.GET("/tests", controllers.GetTests)
	router.POST("/tests", controllers.CreateTest)

	// examples
	// router.POST("/albums", controllers.PostAlbums)

	router.Run("localhost:8080")
}
