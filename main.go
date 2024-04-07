package main

import (
	//	"fmt"
	"net/http"
	"orders_bot/database"

	"github.com/gin-gonic/gin"
)

func main(){
	R := gin.Default()
	R.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
	})
	R.GET("/employees", database.GetEmployees)
	R.GET("/employee/:code", database.GetEmployee)
	R.POST("/employee", database.AddEmployees)
	R.PATCH("/employee",  database.UpdateEmployee)
	database.ConnectDatabase()
	R.Run()
}

