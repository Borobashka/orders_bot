package tgbot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Employee struct {
	Employee_id  int    `json:"employee_id"`
	Name         string `json:"name"`
	Creationdate string `json:"crteationdate"`
	Exhausted    bool   `json:"exhausted"`
	Role         string `json:"role"`
	Phone        string `json:"phone"`
	State        int
}

var CreateEmployee map[int64]*Employee

type ChatBot struct {
	Bot           *tgbotapi.BotAPI
	ChatId        int64
	UpdChannel    tgbotapi.UpdatesChannel
	Update        tgbotapi.Update
	UpadateConfig tgbotapi.UpdateConfig
}
