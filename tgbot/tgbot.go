package tgbot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

)

func StartBot() {

	ChatBot := ChatBot{}

	ChatBot.ConnectDatabase()

	for update := range ChatBot.UpdChannel {
		if update.Message != nil { // If we got a message
			fmt.Printf("from: %s; chatID: %v message: %s\n", 
			update.Message.From.FirstName, 
			update.Message.Chat.ID,
			update.Message.Text) 

			if update.Message.Text == "ping" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "pong")
				msg.ReplyToMessageID = update.Message.MessageID
				ChatBot.Bot.Send(msg)
			}

			if update.Message.Text == "createEmployee" {
				CreateEmployee = make(map[int64]*Employee)
				CreateEmployee[update.Message.From.ID] = new(Employee)
				CreateEmployee[update.Message.From.ID].State = 0
				CreateEmployee[update.Message.From.ID].Employee_id = int(update.Message.Chat.ID)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "введите ваше имя")
				ChatBot.Bot.Send(msg)
			} else { 
				Employee, ok := CreateEmployee[update.Message.Chat.ID]
				if ok{
					if Employee.State == 0 {
						Employee.Name = update.Message.Text
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "введите ваш номер телефона")
						ChatBot.Bot.Send(msg)
						Employee.State = 1  
					} else if Employee.State == 1 {	
						Employee.Phone = update.Message.Text
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "введите вашу роль")
						ChatBot.Bot.Send(msg)
						Employee.State = 2
					}  else if Employee.State == 2 {
						Employee.Role = update.Message.Text
						fmt.Println(Employee)
						createNewEmployee(*Employee)
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пользователь создан")
						ChatBot.Bot.Send(msg)
					}
				}
			}
		}
	}
}