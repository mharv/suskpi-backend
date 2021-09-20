package main

import (
	"fmt"
	"os"
	"strconv"
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
	// router.POST("/bauBenchmarks", controllers.CreateBauBenchmark)

	router.GET("/tests", controllers.GetTests)
	router.POST("/tests", controllers.CreateTest)

	// examples
	// router.POST("/albums", controllers.PostAlbums)

	port := fmt.Sprint(getHTTPPort())

	router.Run("localhost:" + port)
}

func getHTTPPort() int {
	httpPort := 8080
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		httpPort, err := strconv.Atoi(val)
		if err == nil {
			return httpPort
		}
	}
	return httpPort
}
