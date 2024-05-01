package api

import(
		"net/http"
		"orders_bot/database"
		"github.com/gin-gonic/gin"
)

func StartApi(){
	R := gin.Default()
	R.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
	})
	R.GET("/employees", database.GetEmployees)
	R.GET("/employee/:code", database.GetEmployee)
	R.POST("/employee", database.AddEmployee)
	R.PATCH("/employee",  database.UpdateEmployee)
	R.DELETE("/employee/:code", database.DeleteEmployee)

	R.GET("/documents", database.GetDocuments)
	R.GET("/maxiddocuments", database.MaxIdDocument)
	R.GET("/document/:code", database.GetDocument)
	R.POST("/document", database.AddDocument)
	R.PATCH("/document",  database.UpdateDocument)
	R.DELETE("/document/:code", database.DeleteDocument)

	R.Run()
}