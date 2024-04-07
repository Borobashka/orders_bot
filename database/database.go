package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"orders_bot/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB 

func ConnectDatabase(){

	err := godotenv.Load()
	if err != nil{
		logger.Error.Log("Error loading .env file: ", "err")
	} else {
		logger.Info.Log("Successfully read env!", "")
	}

	//read our .env file
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, pass)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		logger.Error.Log("Error connecting to the database: ", "errSql")
		panic(errSql)
	} else {
		Db = db
		logger.Info.Log("Successfully connected!", "")
	}
}

func GetEmployees(c *gin.Context) {

	employees := GettingEmployees()

	if employees == nil || len(employees) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		logger.Error.Log("Error loading GetEmployees", "err")
	} else {
		c.IndentedJSON(http.StatusOK, employees)
		logger.Info.Log("Successfully get all employees", "")
	}
}

func GetEmployee(c *gin.Context) {

	code := c.Param("code")
	product := GettingtEmployee(code)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)
		logger.Error.Log("Error loading GetEmployee", "err")
	} else {
		c.IndentedJSON(http.StatusOK, product)
		logger.Info.Log("Successfully get employee", "")
	}
}

func AddEmployees (c *gin.Context) {

	var empl Employee

	if err := c.BindJSON(&empl); err != nil {
		fmt.Println(empl)
		c.AbortWithStatus(http.StatusBadRequest)
		logger.Error.Log("Error loading AddEmployees", "err")
	} else {
		logger.Info.Log("Successfully start AddEmployees!", "")
		AddingEmployee(empl)
		c.IndentedJSON(http.StatusCreated, empl)
		logger.Info.Log("Successfully add employees", "")
	}
}

func UpdateEmployee(c *gin.Context) {

	var empl Employee

	if err := c.BindJSON(&empl); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		logger.Error.Log("Error loading BindJSON UpdateEmployee ", "err")
	} else {
		logger.Info.Log("Successfully start UpdateEmployee!", "")
		UpEmployee(empl)
		c.IndentedJSON(http.StatusCreated, empl)
		logger.Info.Log("Successfully update employees", "")
	}
}

func DeleteEmployee(c *gin.Context) {

	code := c.Param("code")

	if err := code; err == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		logger.Error.Log("Error loading DeleteEmployee", "err")
	} else {
		logger.Info.Log("Successfully start DeleteEmployee!", "")
		DelEmployee(code)
		c.IndentedJSON(http.StatusCreated, "Employee with id:" + code + " was deleted")
	}
}


