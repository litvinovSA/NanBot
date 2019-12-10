package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"log"
)
type Order struct {
	id int
	orderType string
	productName string
	featureName string
	size string
	customerTgID string
	customerName string
}

func main(){
	fmt.Println("\u2648;")
	bot, err := tgbotapi.NewBotAPI("910932452:AAFUsTTegZxiin7oAPJ-D8AImMPfT1EQ2cE")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	var newOrder Order
	for update := range updates {
		if update.CallbackQuery != nil{
			fmt.Print(update)

			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID,update.CallbackQuery.Data))
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			switch update.CallbackQuery.Data {
			case "blank":
				msg.Text = "Отлично, теперь выбери тип изделия"
				msg.ReplyMarkup = typeKeyboard
			}
			if update.CallbackQuery.Data == "blank"{

				bot.Send(msg)
			}
			bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,update.CallbackQuery.Data))
		}
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			switch update.Message.Text {
			case "new":
				newOrder.id = 1
				msg.Text = "Пошив или Бланк"
				msg.ReplyMarkup = numericKeyboard

			}

			bot.Send(msg)
		}
	}
}
