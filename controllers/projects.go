package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sus-kpi-backend/db"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {

	db.ConnectToDb()

	result, err := db.GetData("dbo.ProjectTable")
	if err != nil {
		log.Fatal("Error reading bau benchmarks: ", err.Error())
	}

	byt := []byte(result)
	var dat []interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, dat)
}
