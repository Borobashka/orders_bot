package tgbot

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"orders_bot/logger"
	"os"

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


func createNewEmployee( employee Employee) {
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