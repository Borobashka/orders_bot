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

// type Crud interface{
// 	Get()
// 	Post()
// 	Update()
// 	Delete()
// }

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
		fmt.Println("Не найдено")
	} else {
		c.IndentedJSON(http.StatusOK, employees)
		fmt.Println("Ваши продукты")
	}
}

func GetEmployee(c *gin.Context) {

	code := c.Param("code")
	product := GettingtEmployee(code)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}

func AddEmployees (c *gin.Context) {

	var empl Employee

	if err := c.BindJSON(&empl); err != nil {
		fmt.Println(empl)
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println("Не найдено")
	} else {
		fmt.Println("Найдено")
		AddingEmployee(empl)
		c.IndentedJSON(http.StatusCreated, empl)
		fmt.Println("Обработанно")
	}
}

func UpdateEmployee(c *gin.Context) {

	var empl Employee

	if err := c.BindJSON(&empl); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println("Не найдено")
	} else {
		fmt.Println("Найдено")
		UpEmployee(empl)
		c.IndentedJSON(http.StatusCreated, empl)
		fmt.Println("Обработанно")
	}
}


