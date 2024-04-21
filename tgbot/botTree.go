package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)



var Str = tgbotapi.NewMessage(1, "Hello")

// func BotFunc(str *tgbotapi.Message, chatBot *ChatBot){ 

// }

func (n *Node)InsertFuncLeft(key int, name string){
	if n.Left == nil{
		n.Left = &Node{Key:n.Key+1, NameFunc:name}
		return
	} else if n.Left != nil{
		n.Left.InsertFuncLeft(key, name)
	}	
}

func (n *Node)InsertFuncRight(key int, name string){
	if n.Right == nil{
		n.Right = &Node{Key:n.Key+1, NameFunc:name}
		return
	} else if n.Right != nil{
		n.Right.InsertFuncRight(key, name)
	}
}

// func (n *Node)SearchFunc(name string) bool{
// 	if n.NameFunc == name{
// 		fmt.Println(n.Key, n.Left, n.Right)
// 		return true
// 	}
// 	if n.Left != nil{
// 		if n.Left.SearchFunc(name){
// 			return true
// 		}
// 	}
// 	if n.Right != nil{
// 		if n.Right.SearchFunc(name){
// 			return true
// 		}
// 	}
// 	return false
// }	

// func CreateTree() *Node{
// 	Tree := &Node{Key:1, NameFunc:"Create", BotFunc:func(str *tgbotapi.Message, chatBot *ChatBot){
// 		fmt.Println(str)}, 
// 		Left:nil, Right:nil}

// 	Tree.InsertFuncLeft(1, "Document", BotFunc)
// 	Tree.InsertFuncRight(1, "Employee", BotFunc)
// 	Tree.InsertFuncLeft(2, "InsertDocName", BotFunc)
// 	Tree.InsertFuncLeft(3, "InsertDocAuthor", BotFunc)
// 	Tree.InsertFuncLeft(4, "InsertDocNumber", BotFunc)
 
//     fmt.Println(Tree.SearchFunc("InsertDocAuthor"))
// 	fmt.Println(Tree.Left)
// 	return Tree
// }



