package main

import (
	"orders_bot/api"
	"orders_bot/database"
	"orders_bot/tgbot"
)

func main(){
	database.ConnectDatabase()
	go api.StartApi()
	tgbot.StartBot()
}

