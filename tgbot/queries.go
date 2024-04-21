package tgbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"orders_bot/logger"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)


type SettingBot interface{
	ConnectDatabase()
}

func (chatBot *ChatBot) ConnectDatabase() ChatBot {
	err := godotenv.Load()
	if err != nil{
		logger.Error.Log("Error loading .env file BOTTOKEN: ", "err")
	} else {
		logger.Info.Log("Successfully read env BOTTOKEN!", "")
	}

	//read our .env file
	tokenBot := os.Getenv("BOTTOKEN")


	bot, err := tgbotapi.NewBotAPI(tokenBot)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	chatBot.Bot = bot

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	chatBot.UpdChannel = updates
	return *chatBot
}


func createNewEmployee(employee Employee) {
	loadBytes, err := json.Marshal(employee)
	if err != nil {
		logger.Error.Log("error: неверная структура", "")
	}

	body := bytes.NewReader(loadBytes)

	req, err := http.NewRequest("POST", "http://localhost:8080/employee", body)
	if err != nil {
		logger.Error.Log("error: неверный запрос", "")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error.Log("error: ", "")
	}
	defer resp.Body.Close()
}

func createNewDocument(document Document) {
	loadBytes, err := json.Marshal(document)
	if err != nil {
		logger.Error.Log("error: неверная структура", "")
	}

	body := bytes.NewReader(loadBytes)

	req, err := http.NewRequest("POST", "http://localhost:8080/document", body)
	if err != nil {
		logger.Error.Log("error: неверный запрос", "")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error.Log("error: ", "")
	}
	defer resp.Body.Close()
}

func selectEmployees( bot *tgbotapi.BotAPI, chatId int64){
	

	resp, err := http.Get("http://localhost:8080/employees")
	if err != nil {
		logger.Error.Log("error: неверный запрос", "")
	}

	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		logger.Error.Log("error: ", "")
	}

	var empl []Employee
	err = json.Unmarshal(body, &empl)
	if err != nil {
		logger.Error.Log("error: неверная структура", "")
	}
	res := ""
	for i := range empl {
		Id := empl[i].Employee_id
		Name := empl[i].Name
		CreateDate := empl[i].Creationdate
		Exhausted := empl[i].Exhausted
		Role := empl[i].Role
		Phone := empl[i].Phone
		res = "Id пользователя:" + strconv.Itoa(Id) + "\n" + "Имя пользователя:" + Name + "\n" + "Дата создания:" + CreateDate + "\n" + "Роль:" + Role + "\n" + "Телефон:" + Phone + "\n" + "Статус:" + strconv.FormatBool(Exhausted)
		bot.Send(tgbotapi.NewMessage(chatId,res))
		fmt.Println("Ваш список сотрудников: ",res)
	}

	fmt.Println("Ваш список сотрудников: ",json.Unmarshal(body, &empl))
}

func selectDocuments( bot *tgbotapi.BotAPI, chatId int64){
	
	resp, err := http.Get("http://localhost:8080/documents")
	if err != nil {
		logger.Error.Log("error: неверный запрос", "")
	}

	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		logger.Error.Log("error: ", "")
	}

	var doc []Document
	err = json.Unmarshal(body, &doc)
	if err != nil {
		logger.Error.Log("error: неверная структура", "")
	}
	res := ""
	for i := range doc {
		Document_id := doc[i].Document_id
		Year := doc[i].Year
		Name := doc[i].Name
		Author := doc[i].Author
		Creationdate := doc[i].Creationdate

		res = "Номер приказа:" + strconv.Itoa(Document_id) + "\n" + "Год:" + strconv.Itoa(Year) + "\n" + "Название:" + Name + "\n" + "Автор:" + Author + "\n" + "Дата создания:" + Creationdate
		bot.Send(tgbotapi.NewMessage(chatId,res))
		fmt.Println("Ваш список приказов: ",res)
	}
}

func selectDocument( bot *tgbotapi.BotAPI, chatId int64, code string){

	resp, err := http.Get("http://localhost:8080/document/" + code)
	if err != nil {
		logger.Error.Log("error: неверный запрос", "")
	}

	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		logger.Error.Log("error: ", "")
	}

	var doc Document
	err = json.Unmarshal(body, &doc)
	if err != nil {
		logger.Error.Log("error: неверная структура", "")
	}
	Document_id := doc.Document_id
	Year := doc.Year
	Name := doc.Name
	Author := doc.Author
	Creationdate := doc.Creationdate

	res := "Номер приказа:" + strconv.Itoa(Document_id) + "\n" + "Год:" + strconv.Itoa(Year) + "\n" + "Название:" + Name + "\n" + "Автор:" + Author + "\n" + "Дата создания:" + Creationdate
	bot.Send(tgbotapi.NewMessage(chatId,res))
	fmt.Println("Ваш приказ: ",res)
}