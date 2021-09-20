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

func GetProjects(c *gin.Context) {

	db.ConnectToDb()

	result, err := db.GetData("dbo.ProjectTable")
	if err != nil {
		log.Fatal("Error reading projects: ", err.Error())
	}

	byt := []byte(result)
	var dat []interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, dat)
}

func GetProjectById(c *gin.Context) {
	id := c.Param("id")

	db.ConnectToDb()

	result, err := db.GetDataById("dbo.ProjectTable", id)
	if err != nil {
		log.Fatal("Error reading single project: ", err.Error())
	}

	byt := []byte(result)
	var dat []interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, dat)
}

func CreateProject(c *gin.Context) {

	db.ConnectToDb()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	response, err := db.PostNewItem("dbo.ProjectTable", jsonData)

	var dat map[string]interface{}

	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, dat)

}
