package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	_ "github.com/jackc/pgx/stdlib"
	"log"
	"strings"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("910932452:AAFUsTTegZxiin7oAPJ-D8AImMPfT1EQ2cE")
	if err != nil {
		log.Panic(err)
	}
	db := initConnection()
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
					newOrder = initOrder(update.Message.From.UserName, getNextID(db))
					orders[update.Message.Chat.ID] = newOrder
				}
			}
			id = update.Message.Chat.ID
		} else {
			id = update.CallbackQuery.Message.Chat.ID
		}
		if (update.Message != nil && !update.Message.IsCommand()) || update.CallbackQuery != nil {
			if (update.Message != nil &&
				(update.Message.Chat.UserName == "" ||
					update.Message.Chat.UserName == "")) ||
				(update.CallbackQuery != nil &&
					(update.CallbackQuery.From.UserName == "" ||
						update.CallbackQuery.From.UserName == "")) {
				adminServe(bot, update, id, db)
			} else {
				msg := NewServe(update, orders[id], id, db)
				_, err = bot.Send(msg)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

//func main() {
//	db := initConnection()
//	order := Order{
//		Orderid:      uuid.New(),
//		Type:        "Blank",
//		ProductName: "T-shirt",
//		Features:    []string{"memes", "kekes"},
//		Amount:      100,
//		Cols:        5,
//		Mockup:      "memkek",
//		Layout:      "memkek",
//		CustomerID:  "1488",
//		Deadline:    "10 april",
//		Comment:     "biba i boba",
//		State:       "new",
//		edit:        false,
//	}
//	putOrder(order, db)
//	getNewOrders(db)
//	db.Close()
//}
