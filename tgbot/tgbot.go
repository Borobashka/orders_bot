package tgbot

import (
	"fmt"
	"orders_bot/logger"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


var createEmployee = make(map[int64]*Employee)

var createDocument = make(map[int64]*Document)


type Node struct{
	Key int
    NameFunc string
	BotFunc func(str string, Tree *Node)
	State bool
	Left *Node
    Right *Node
}


var MainMenu = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
        tgbotapi.NewKeyboardButton("получить"),
        tgbotapi.NewKeyboardButton("создать"),
        tgbotapi.NewKeyboardButton("обновить"),
    ),
    // tgbotapi.NewKeyboardButtonRow(
    //     tgbotapi.NewKeyboardButton("4"),
    //     tgbotapi.NewKeyboardButton("5"),
    //     tgbotapi.NewKeyboardButton("6"),
    // ),
)

var OrderOrEmployee = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
        tgbotapi.NewKeyboardButton("документ"),
        tgbotapi.NewKeyboardButton("пользователя"),
    ),
)

var GetOrderOrEmployee = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
        tgbotapi.NewKeyboardButton("один документ"),
		tgbotapi.NewKeyboardButton("все документы"),
        tgbotapi.NewKeyboardButton("всех пользователей"),
    ),
)



func StartBot() {
	ChatBot := &ChatBot{}
	ChatBot.ConnectDatabase()

	var (
		id int
		UserMessage string
	)

	Tree := &Node{
		Key: id, 
		NameFunc: "", 
		BotFunc: func(str string, Tree *Node){},
		State: true,
		Left:nil, 
		Right:nil,
	}

	for update := range ChatBot.UpdChannel {
		
		ChatBot.Update = update

		if message := update.Message; message != nil {
			UserMessage = message.Text
		}

		if UserMessage == "start" {
			msg := tgbotapi.NewMessage(ChatBot.Update.Message.Chat.ID, "Привет, жми на кнопку")
			msg.ReplyToMessageID = ChatBot.Update.Message.MessageID
			msg.ReplyMarkup = MainMenu
			ChatBot.Bot.Send(msg)
		}


		//create
		if UserMessage == MainMenu.Keyboard[0][1].Text {
			Tree.Key = 1
			Tree.NameFunc = "create"
			Tree.BotFunc = func(str string, Tree *Node){
				msg := tgbotapi.NewMessage(ChatBot.Update.Message.Chat.ID, "Вы хотите создать документ или пользователя?")
				msg.ReplyToMessageID = ChatBot.Update.Message.MessageID
				msg.ReplyMarkup = OrderOrEmployee
				ChatBot.Bot.Send(msg)
			}
			Tree.BotFunc(Tree.NameFunc, Tree)
		} else if Tree.NameFunc == "create" && UserMessage == OrderOrEmployee.Keyboard[0][1].Text {
			Tree.InsertFuncLeft(Tree.Key, "employee")
			Tree.Left.BotFunc = func(str string, Tree *Node){				
				CreateEmployee = make(map[int64]*Employee)
				CreateEmployee[update.Message.From.ID] = new(Employee)
				CreateEmployee[update.Message.From.ID].State = 0
				CreateEmployee[update.Message.From.ID].Employee_id = int(update.Message.Chat.ID)

				msg := tgbotapi.NewMessage(ChatBot.Update.Message.Chat.ID, "введите ваше имя")
				msg.ReplyToMessageID = ChatBot.Update.Message.MessageID
				ChatBot.Bot.Send(msg)
			}
			Tree.Left.BotFunc(Tree.NameFunc, Tree)
			Tree.Left.NameFunc = "employee"
		} else if Tree.NameFunc == "create" && Tree.Left != nil && Tree.Left.NameFunc == "employee" && UserMessage != "employee" {
			Employee := CreateEmployee[update.Message.From.ID]
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
				msg.ReplyMarkup = MainMenu
				ChatBot.Bot.Send(msg)
				Employee.State = 0
				Tree.Left.NameFunc = ""
				Tree.NameFunc = ""
			}
		} else if Tree.NameFunc == "create" && UserMessage == OrderOrEmployee.Keyboard[0][0].Text {
			Tree.InsertFuncRight(Tree.Key, "document")
			Tree.Right.BotFunc = func(str string, Tree *Node){
				CreateDocument = make(map[int64]*Document)
				CreateDocument[update.Message.From.ID] = new(Document)
				CreateDocument[update.Message.From.ID].State = 0

				msg := tgbotapi.NewMessage(ChatBot.Update.Message.Chat.ID, "введите год создание документа")
				msg.ReplyToMessageID = ChatBot.Update.Message.MessageID
				ChatBot.Bot.Send(msg)
			}
			Tree.Right.BotFunc(Tree.NameFunc, Tree)
			Tree.Right.NameFunc = "document"
		} else if Tree.NameFunc == "create" && Tree.Right != nil && Tree.Right.NameFunc == "document" && UserMessage != "document" {
			Document := CreateDocument[update.Message.From.ID]
			if Document.State == 0 {
				year,_ := strconv.Atoi(update.Message.Text)
				Document.Year = year
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "введите название документа")
				ChatBot.Bot.Send(msg)
				Document.State = 1
			} else if Document.State == 1 {
				Document.Name = update.Message.Text
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "введите имя автора")
				ChatBot.Bot.Send(msg)
				Document.State = 2
			} else if Document.State == 2 {
				Document.Author = update.Message.Text
				Document.Employee_id = int(update.Message.Chat.ID)
				fmt.Println(Document)
				createNewDocument(*Document)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Документ создан")
				msg.ReplyMarkup = MainMenu
				ChatBot.Bot.Send(msg)
				pandocdoc := PandocCreateDocument()
				if pandocdoc != nil{
					logger.Error.Log("error: ", "")
				}
				Document.State = 0
				Tree.Right.NameFunc = ""
				Tree.NameFunc = ""
			} 
		}

		//get
		if UserMessage == MainMenu.Keyboard[0][0].Text  {
			Tree.Key = 1 
			Tree.NameFunc = "get"
			Tree.BotFunc = func(str string, Tree *Node) {
				msg := tgbotapi.NewMessage(ChatBot.Update.Message.Chat.ID, "Что вы хотите получить?")
				msg.ReplyToMessageID = ChatBot.Update.Message.MessageID
				msg.ReplyMarkup = GetOrderOrEmployee
				ChatBot.Bot.Send(msg)
			}
			Tree.BotFunc(Tree.NameFunc, Tree)
		} else if Tree.NameFunc == "get" && UserMessage == GetOrderOrEmployee.Keyboard[0][0].Text {
			Tree.InsertFuncLeft(Tree.Key, "oneDocument")
			Tree.Left.BotFunc = func(str string, Tree *Node){
				msg := tgbotapi.NewMessage(ChatBot.Update.Message.Chat.ID, "введите id документа")
				msg.ReplyToMessageID = ChatBot.Update.Message.MessageID
				ChatBot.Bot.Send(msg)
			}
			Tree.Left.BotFunc(Tree.NameFunc, Tree)
		} else if Tree.NameFunc == "get" && Tree.Left != nil && Tree.Left.NameFunc == "oneDocument" && UserMessage != "oneDocument" {
			msg := tgbotapi.NewMessage(ChatBot.Update.Message.Chat.ID, "Ваш документ")
			msg.ReplyToMessageID = ChatBot.Update.Message.MessageID
			msg.ReplyMarkup = MainMenu
			ChatBot.Bot.Send(msg)
			selectDocument(ChatBot.Bot, update.Message.Chat.ID, update.Message.Text)
			Tree.Left = nil
			Tree.NameFunc = ""
		} else if Tree.NameFunc == "get" && UserMessage ==  GetOrderOrEmployee.Keyboard[0][2].Text {
			msg := tgbotapi.NewMessage(ChatBot.Update.Message.Chat.ID, "Ваши сотрудники")
			msg.ReplyToMessageID = ChatBot.Update.Message.MessageID
			msg.ReplyMarkup = MainMenu
			ChatBot.Bot.Send(msg)
			selectEmployees(ChatBot.Bot, update.Message.Chat.ID)
			Tree.NameFunc = ""
		} else if Tree.NameFunc == "get" && UserMessage ==  GetOrderOrEmployee.Keyboard[0][1].Text {
			msg := tgbotapi.NewMessage(ChatBot.Update.Message.Chat.ID, "Ваши документы")
			msg.ReplyToMessageID = ChatBot.Update.Message.MessageID
			msg.ReplyMarkup = MainMenu
			ChatBot.Bot.Send(msg)
			selectDocuments(ChatBot.Bot, update.Message.Chat.ID)

			Tree.NameFunc = ""
		} 

		// update
		if UserMessage == MainMenu.Keyboard[0][2].Text  {
			Tree.Key = 1
			Tree.NameFunc = "update"
			Tree.BotFunc = func(str string, Tree *Node){
				msg := tgbotapi.NewMessage(ChatBot.Update.Message.Chat.ID, "Вы хотите обновить документ или пользователя?")
				msg.ReplyToMessageID = ChatBot.Update.Message.MessageID
				msg.ReplyMarkup = OrderOrEmployee
				ChatBot.Bot.Send(msg)
			}
			Tree.BotFunc(Tree.NameFunc, Tree)
		} else if Tree.NameFunc == "update" && UserMessage == OrderOrEmployee.Keyboard[0][1].Text {
			Tree.InsertFuncLeft(Tree.Key, "employee")
			Tree.Left.BotFunc = func(str string, Tree *Node){				
				CreateEmployee = make(map[int64]*Employee)
				CreateEmployee[update.Message.From.ID] = new(Employee)
				CreateEmployee[update.Message.From.ID].State = 0
				CreateEmployee[update.Message.From.ID].Employee_id = int(update.Message.Chat.ID)

				msg := tgbotapi.NewMessage(ChatBot.Update.Message.Chat.ID, "введите id сотрудника")
				msg.ReplyToMessageID = ChatBot.Update.Message.MessageID
				ChatBot.Bot.Send(msg)
			}
			Tree.Left.BotFunc(Tree.NameFunc, Tree)
			Tree.Left.NameFunc = "employee"
		} else if Tree.NameFunc == "update" && Tree.Left != nil && Tree.Left.NameFunc == "employee" && UserMessage != "employee" {
			Employee := CreateEmployee[update.Message.From.ID]
			if Employee.State == 0 {
				id, _ :=  strconv.Atoi(update.Message.Text)
				Employee.Employee_id = id
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "введите новое имя сотрудника или введите то, которое было")
				ChatBot.Bot.Send(msg)
				Employee.State = 1  
			} else if Employee.State == 1 {	
				Employee.Name = update.Message.Text
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "введите роль сотрудника или введите ту, которая была")
				ChatBot.Bot.Send(msg)
				Employee.State = 2
			}  else if Employee.State == 2 {
				Employee.Role = update.Message.Text
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "введите телефон сотрудника или введите тот, который был")
				ChatBot.Bot.Send(msg)
				Employee.State = 3
			} else if Employee.State == 3 {
				Employee.Phone = update.Message.Text
				fmt.Println(Employee)
				updateEmployee(*Employee)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пользователь обвновлен")
				msg.ReplyMarkup = MainMenu
				ChatBot.Bot.Send(msg)
				Employee.State = 0
				Tree.Left.NameFunc = ""
				Tree.NameFunc = ""
			}
		} else if Tree.NameFunc == "update" && UserMessage == OrderOrEmployee.Keyboard[0][0].Text {
			Tree.InsertFuncRight(Tree.Key, "document")
			Tree.Right.BotFunc = func(str string, Tree *Node){
				CreateDocument = make(map[int64]*Document)
				CreateDocument[update.Message.From.ID] = new(Document)
				CreateDocument[update.Message.From.ID].State = 0

				msg := tgbotapi.NewMessage(ChatBot.Update.Message.Chat.ID, "введите номер документа, который хотите обновить")
				msg.ReplyToMessageID = ChatBot.Update.Message.MessageID
				ChatBot.Bot.Send(msg)
			}
			Tree.Right.NameFunc = "document"
			Tree.Right.BotFunc(Tree.NameFunc, Tree)
		} else if Tree.NameFunc == "update" && Tree.Right != nil && Tree.Right.NameFunc == "document" && UserMessage != "document" {
			Document := CreateDocument[update.Message.From.ID]
			if Document.State == 0 {
				id, _ :=  strconv.Atoi(update.Message.Text)
				Document.Document_id = id
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "введите год документа")
				ChatBot.Bot.Send(msg)
				Document.State = 1
			} else if Document.State == 1 {	
				year,_ := strconv.Atoi(update.Message.Text)
				Document.Year = year
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "введите название документа")
				ChatBot.Bot.Send(msg)
				Document.State = 2
			} else if Document.State == 2 {
				Document.Name = update.Message.Text
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "введите имя автора")
				ChatBot.Bot.Send(msg)
				Document.State = 3
			} else if Document.State == 3 {
				Document.Author = update.Message.Text
				Document.Employee_id = int(update.Message.Chat.ID)
				fmt.Println(Document)
				updateDocument(*Document)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Документ обновлен")
				msg.ReplyMarkup = MainMenu
				ChatBot.Bot.Send(msg)
				Document.State = 0
				Tree.Right.NameFunc = ""
				Tree.NameFunc = ""
			} 
		}
	}
}



