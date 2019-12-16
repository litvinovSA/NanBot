package main

import (
	"github.com/google/uuid"
)

type Order struct {
	id          uuid.UUID      `db: orderid`
	Type        string   `db: type`
	ProductName string   `db: productname`
	Features    []string `db: Features`
	Amount      int      `db: Amount`
	Cols        int      `db: cols`
	Mockup      string   `db: Mockup`
	Layout      string   `db: Layout`
	CustomerID  string   `db: customerid`
	Deadline    string   `db: Deadline`
	Comment     string   `db: Comment`
	State       string   `db: State`
	edit        bool
}

var orders = make(map[int64]*Order)

func initOrder() *Order {
	return &Order{
		id:          uuid.New(),
		Type:        "",
		ProductName: "",
		Features:    nil,
		Amount:      0,
		Cols:        0,
		Mockup:      "",
		Layout:      "",
		CustomerID:  "",
		Deadline:    "",
		Comment:     "",
		State:       "new",
		edit:        false,
	}
}

//func main() {
//	bot, err := tgbotapi.NewBotAPI("910932452:AAFUsTTegZxiin7oAPJ-D8AImMPfT1EQ2cE")
//	if err != nil {
//		log.Panic(err)
//	}
//	bot.Debug = true
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//	u := tgbotapi.NewUpdate(0)
//	u.Timeout = 60
//	updates, err := bot.GetUpdatesChan(u)
//	var newOrder *Order
//	for update := range updates {
//		var id int64
//		if update.Message != nil {
//			if update.Message.Text != "" {
//				if strings.ToLower(update.Message.Text) == "new" {
//					newOrder = initOrder()
//					orders[update.Message.Chat.ID] = newOrder
//
//				}
//			}
//			id = update.Message.Chat.ID
//		} else {
//			id = update.CallbackQuery.Message.Chat.ID
//		}
//		if (update.Message != nil && !update.Message.IsCommand()) || update.CallbackQuery !=nil {
//			msg := Serve(update, orders[id], id)
//			bot.Send(msg)
//		}
//	}
//}

func main() {
	db := initConnection()
	order  := Order{
		id:          uuid.New(),
		Type:        "Blank",
		ProductName: "T-shirt",
		Features:    nil,
		Amount:      100,
		Cols:        5,
		Mockup:      "memkek",
		Layout:      "memkek",
		CustomerID:  "@mrdken",
		Deadline:    "10 april",
		Comment:     "biba i boba",
		State:       "new",
		edit:        false,
	}
	putOrder(order, db)
	db.Close()
}
