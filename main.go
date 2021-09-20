package main

import (
	"fmt"
	"os"
	"strconv"
	"sus-kpi-backend/controllers"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/gin-gonic/gin"
)

// see below link for building unix exe on windows
// https://medium.com/@utranand/building-golang-package-for-linux-from-windows-22fa23764808

func main() {

	// Routes
	router := gin.Default()
	router.GET("/bauBenchmarks", controllers.GetBauBenchmarks)
	router.GET("/projects", controllers.GetProjects)
	router.GET("/bauBenchmarks/:id", controllers.GetBauBenchmarkById)
	router.GET("/projects/:id", controllers.GetProjectById)
	router.POST("/projects", controllers.CreateProject)
	router.POST("/bauBenchmarks", controllers.CreateBauBenchmark)

	// testing routes
	router.GET("/tests", controllers.GetTests)
	router.POST("/tests", controllers.CreateTest)

	// get port if on azure
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
