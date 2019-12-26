package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	_ "github.com/jackc/pgx/stdlib"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
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
			id = update.Message.Chat.ID
			if update.Message.IsCommand() {
				switch update.Message.Text {
				case "/start":
					msg := tgbotapi.NewMessage(id, "Привет, "+update.Message.From.FirstName+".\n Чем мы займемся сегодня?")
					msg.ReplyMarkup = startKeyboard
					_, err = bot.Send(msg)
					if err != nil {
						log.Println(err)
					}
					continue
				}
			}
			if update.Message.Text != "" {
				if update.Message.Text == l10n["Contacts"] {
					msg := tgbotapi.NewMessage(id, l10n["contactsData"])
					_, err := bot.Send(msg)
					if err != nil {
						log.Fatal(err)
					}
					continue
				} else if update.Message.Text == l10n["Orders"] {
					msg := tgbotapi.NewMessage(id, "")
					orders := getUserOrders(db, update.Message.From.UserName)
					for _, order := range orders {
						msg.Text = stringifyOrder(&order) + "\n Статус: " + l10n[order.State]
						_, err := bot.Send(msg)
						if err != nil {
							log.Fatal(err)
						}
					}
					continue
				} else if update.Message.Text == l10n["NewOrder"] || update.Message.Text == l10n["another"] {
					newOrder = initOrder(update.Message.From.UserName, getNextID(db))
					orders[update.Message.Chat.ID] = newOrder
				}
			}
		} else {
			id = update.CallbackQuery.Message.Chat.ID
		}
		if (update.Message != nil && !update.Message.IsCommand()) || update.CallbackQuery != nil {
			if (update.Message != nil &&
				(update.Message.From.ID == 910932452 ||
					update.Message.From.ID == 1)) ||
				(update.CallbackQuery != nil &&
					(update.CallbackQuery.From.ID == 910932452 ||
						update.CallbackQuery.From.ID == 1)) {
				adminServe(bot, update, id, db)
			} else {
				msg := NewServe(update, orders[id], id, db)
				_, err = bot.Send(msg)
				if err != nil {
					log.Fatal(err)
				}
				if orders[id].state == stateFin{
					msg = tgbotapi.NewMessage(id, "")
					msg.Text = getPhoto(orders[id].Mockup)
					bot.Send(msg)
					msg = tgbotapi.NewMessage(id, "")
					msg.Text = getPhoto(orders[id].Layout)
					bot.Send(msg)
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
