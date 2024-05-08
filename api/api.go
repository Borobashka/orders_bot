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
	R.GET("/documents/:year", database.GetDocumentsYear) 
	R.GET("/maxiddocuments/:year", database.MaxIdDocument) // дописать запрос в боте
	R.GET("/document/:code/:year", database.GetDocument) // дописать запрос в боте
	R.POST("/document", database.AddDocument)
	R.PATCH("/document",  database.UpdateDocument)
	R.DELETE("/document/:code/:year", database.DeleteDocument)

	R.Run(":8080")
}