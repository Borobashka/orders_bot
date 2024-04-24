package tgbot

import (	
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var CreateEmployee map[int64]*Employee

var CreateDocument map[int64]*Document

type Employee struct {
	Employee_id  int    `json:"employee_id"`
	Name         string `json:"name"`
	Creationdate string `json:"crteationdate"`
	Exhausted    bool   `json:"exhausted"`
	Role         string `json:"role"`
	Phone        string `json:"phone"`
	State        int
}

type Document struct {
	Document_id 	int 	`json:"document_id"`  //не надо
	Year 			int 	`json:"year"`		  
	Path 			string  `json:"path"`		  //не надо
	Name 			string  `json:"name"`		  
	Author 			string  `json:"author"`
	Creationdate 	string  `json:"creationdate"` //не надо
	Employee_id 	int 	`json:"employee_id"`  //не надо 
	State 			int
}

type ChatBot struct {
	Bot           *tgbotapi.BotAPI
	ChatId        int64
	UpdChannel    tgbotapi.UpdatesChannel
	Update        tgbotapi.Update
	UpadateConfig tgbotapi.UpdateConfig
	FuncNow		  string
	NextStep 	  string	
}
