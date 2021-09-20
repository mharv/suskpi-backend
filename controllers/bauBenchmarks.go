package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sus-kpi-backend/db"

	"github.com/gin-gonic/gin"
)

func GetBauBenchmarks(c *gin.Context) {

	db.ConnectToDb()

	result, err := db.GetData("dbo.BauTable")
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

func GetBauBenchmarkById(c *gin.Context) {
	id := c.Param("id")

	db.ConnectToDb()

	result, err := db.GetDataById("dbo.BauTable", id)
	if err != nil {
		log.Fatal("Error reading single Bau benchmark: ", err.Error())
	}

	byt := []byte(result)
	var dat []interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, dat)
}

func CreateBauBenchmark(c *gin.Context) {

	db.ConnectToDb()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	response, err := db.PostNewItem("dbo.BauTable", jsonData)

	var dat map[string]interface{}

	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, dat)

}
