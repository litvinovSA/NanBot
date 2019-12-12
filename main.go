package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/segmentio/ksuid"
	"log"
	"strings"
)

type Order struct {
	id           ksuid.KSUID
	orderType    string
	productName  string
	features     []string
	size         string
	amount       int
	amountCol    int
	mockup       string
	layout       string
	customerTgID string
	customerName string
	deadline     string
	comment      string
	edit         bool
}

var orders = make(map[int64]*Order)

func initOrder() *Order {
	return &Order{
		id:           ksuid.New(),
		orderType:    "",
		productName:  "",
		features:     nil,
		size:         "",
		amount:       0,
		amountCol:    0,
		mockup:       "",
		layout:       "",
		customerTgID: "",
		customerName: "",
		deadline:     "",
		comment:      "",
		edit:         false,
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI("910932452:AAFUsTTegZxiin7oAPJ-D8AImMPfT1EQ2cE")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	var newOrder *Order
	for update := range updates {
		var id int64
		if update.Message != nil {
			if update.Message.Text != "" {
				if strings.ToLower(update.Message.Text) == "new" {
					newOrder = initOrder()
					orders[update.Message.Chat.ID] = newOrder

				}
			}
			id = update.Message.Chat.ID
		} else {
			id = update.CallbackQuery.Message.Chat.ID
		}
		if update.Message != nil && !update.Message.IsCommand(){
			msg := Serve(update, orders[id], id)
			bot.Send(msg)
		}
	}
}
