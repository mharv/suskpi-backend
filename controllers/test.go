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

func GetTests(c *gin.Context) {

	db.ConnectToDb()

	result, err := db.GetData("dbo.TestTable")
	if err != nil {
		log.Fatal("Error reading tests: ", err.Error())
	}

	byt := []byte(result)
	var dat []interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, dat)
}

func CreateTest(c *gin.Context) {

	db.ConnectToDb()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	response, err := db.PostNewItem("dbo.TestTable", jsonData)

	var dat map[string]interface{}

	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, dat)

}
