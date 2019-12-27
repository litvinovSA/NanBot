package main

import (
	"log"
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

//func Server(msg *tb.Message, bot tb.Bot) {
//	newOrder := orders[msg.Sender.ID]
//	if (msg.Text != l10n["NewOrder"] || msg.Text != l10n["another"]) && newOrder == nil {
//		reply := tb.Message{}
//		reply.Text = "Привет, " + msg.Sender.FirstName + ".\n Чем мы займемся сегодня?"
//		_, err := bot.Send(msg.Sender, reply)
//		if err != nil {
//			log.Println(err)
//		}
//	}
//	switch msg.Text {
//	case l10n["NewOrder"], l10n["another"]:
//		newOrder = initOrder(msg.Sender.Username, getNextID(db))
//	}
//	bot.Send(msg.Sender, getKeyboardAndTextByState(newOrder, msg.Sender.ID))
//
//}

func photoHandler(msg *tb.Message) {

}

func getOrder (id int, username string) *Order{
	newOrder := orders[id]
	if newOrder == nil {
		newOrder = initOrder(username, getNextID(db))
	}
	return newOrder
}

func main() {
	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	bot.Handle("/start", func(msg *tb.Message) {
		text := "Привет, " + msg.Sender.FirstName + ".\n Чем мы займемся сегодня?"
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{ReplyKeyboard:startKeyboard, ResizeReplyKeyboard:true})
		if err != nil {
			log.Println(err)
		}
	})
	// Block 1 hello pane
	bot.Handle(&contactsBtn, func(msg *tb.Message) {
		text := l10n["contactsData"]
		_, err := bot.Send(msg.Sender, text)
		if err != nil {
			log.Println(err)
		}
	})
	bot.Handle(&myordersBtn, func(msg *tb.Message) {
		orders := getUserOrders(msg.Sender.Username)
		for _, order := range orders {
			text := stringifyOrder(&order) + "\n Статус: " + l10n[order.State]
			_, err := bot.Send(msg.Sender, text)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
	bot.Handle(&newOrderBtn, func(msg *tb.Message){
		orders[msg.Sender.ID] = initOrder(msg.Sender.Username, getNextID(db))
		text := steps[stateProdtype]
		orders[msg.Sender.ID].state = stateProdtype
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{ReplyKeyboard:orderTypeKeyboard, ResizeReplyKeyboard:true})
		if err != nil {
			log.Fatal(err)
		}
	})
	// Block 2 order type
	bot.Handle(&sewingBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.Type = msg.Text
		newOrder.state = stateProduct
		text := steps[stateProduct]
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:typeKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	bot.Handle(&blankBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.Type = msg.Text
		newOrder.state = stateProduct
		text := steps[stateCols]
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:typeKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	// Block 3 product
	bot.Handle(&TshirtBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.ProductName = msg.Text
		if newOrder.Type == l10n["Blank"]{
			newOrder.state = stateCols
			text := steps[stateCols]
			_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
				ReplyKeyboard:colorsKeyboard,
				ResizeReplyKeyboard:true,
				OneTimeKeyboard: true,
			})
			if err != nil {
				log.Fatal(err)
			}
		}else {
			text := steps[stateFeature1]
			newOrder.state = stateFeature1
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:tshirtKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard: true,
		})
		if err != nil {
			log.Fatal(err)
		}}
	})
	bot.Handle(&HoodBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.ProductName = msg.Text
		if newOrder.Type == l10n["Blank"]{
			newOrder.state = stateCols
			text := steps[stateCols]
			_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
				ReplyKeyboard:colorsKeyboard,
				ResizeReplyKeyboard:true,
				OneTimeKeyboard: true,
			})
			if err != nil {
				log.Fatal(err)
			}
		}else {
			newOrder.state = stateFeature2
			text := steps[stateFeature2]
			_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
				ReplyKeyboard:hoodKeyboard,
				ResizeReplyKeyboard:true,
				OneTimeKeyboard: true,
			})
			if err != nil {
				log.Fatal(err)
			}}
	})
	bot.Handle(&SweetBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.ProductName = msg.Text
		if newOrder.Type == l10n["Blank"]{
			newOrder.state = stateCols
			text := steps[stateCols]
			_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
				ReplyKeyboard:colorsKeyboard,
				ResizeReplyKeyboard:true,
				OneTimeKeyboard: true,
			})
			if err != nil {
				log.Fatal(err)
			}
		}else {
			newOrder.state = stateFeature1
			text := steps[stateFeature1]
			_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
				ReplyKeyboard:sweetKeybord,
				ResizeReplyKeyboard:true,
				OneTimeKeyboard: true,
			})
			if err != nil {
				log.Fatal(err)
			}}
	})
	// Block 4 feature 1
	bot.Handle(&defaultTShirtBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.Features += msg.Text
		newOrder.state = stateCols
		text := steps[stateCols]
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:colorsKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	bot.Handle(&oversizeTShirtBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.Features += msg.Text
		newOrder.state = stateCols
		text := steps[stateCols]
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:colorsKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	bot.Handle(&SetinBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.Features += msg.Text
		newOrder.state = stateCols
		text := steps[stateCols]
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:colorsKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	bot.Handle(&sweetRelganBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.Features += msg.Text
		newOrder.state = stateCols
		text := steps[stateCols]
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:colorsKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	bot.Handle(&pocketSewingBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.Features += msg.Text
		newOrder.state = stateCols
		text := steps[stateCols]
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:colorsKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	bot.Handle(&pocketSetinBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.Features += msg.Text
		newOrder.state = stateCols
		text := steps[stateCols]
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:colorsKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	// Block 4.1 feature 2
	bot.Handle(&hoodDefault, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.Features += msg.Text + ", "
		newOrder.state = stateFeature1
		text := steps[stateFeature1]
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:pocketKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	bot.Handle(&hoodReglan, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.Features += msg.Text + ", "
		newOrder.state = stateFeature1
		text := steps[stateFeature1]
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:pocketKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	bot.Handle(&hoodOversize, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.Features += msg.Text + ", "
		newOrder.state = stateFeature1
		text := steps[stateFeature1]
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:pocketKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
	// Block 5 Cols
	bot.Handle(&dunnoBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.state = 0
		newOrder.state = stateAmount
		text := steps[stateAmount]
		_, err := bot.Send(msg.Sender, text)
		if err != nil {
			log.Fatal(err)
		}
	})
	bot.Handle(&jpegBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.state = 0
		newOrder.state = stateAmount
		text := steps[stateAmount]
		_, err := bot.Send(msg.Sender, text)
		if err != nil {
			log.Fatal(err)
		}
	})
	// Block 6 Cols, Commentary, Date
	bot.Handle(tb.OnText, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		var text string
		if num, err := strconv.Atoi(msg.Text); err == nil {
			if newOrder.state == stateCols {
				newOrder.Cols = num
				newOrder.state = stateAmount
				text = steps[stateAmount]
				_, err = bot.Send(msg.Sender, text)
				if err != nil {
				    log.Println(err)
				}
			} else if newOrder.state == stateAmount {
				newOrder.Amount = num
				text = steps[stateLayout]
				newOrder.state = stateLayout
				_, err = bot.Send(msg.Sender, text)
				if err != nil {
				    log.Println(err)
				}
			}
		} else {
			if newOrder.state == stateDeadline {
				newOrder.Deadline = msg.Text
				newOrder.state = stateComment
				text = steps[stateComment]
				_, err = bot.Send(msg.Sender, text)
				if err != nil {
				    log.Println(err)
				}
			} else if newOrder.state == stateComment{
				newOrder.Comment = msg.Text
				newOrder.state = stateFin
				text = steps[stateFin]
				text += "\n" + stringifyOrder(newOrder)
				_, err = bot.Send(msg.Sender, text, &tb.ReplyMarkup{
					ReplyKeyboard:finKeyboard,
					ResizeReplyKeyboard:true,
					OneTimeKeyboard: true,
				})
				if err != nil {
				    log.Println(err)
				}
			} else if newOrder.state == stateLayout {
				if isValidUrl(msg.Text) {
					newOrder.Layout = msg.Text
					newOrder.state = stateMock
					text = steps[stateMock]
					_, err = bot.Send(msg.Sender, text)
					if err != nil {
						log.Println(err)
					}
				}
			} else if newOrder.state == stateMock{
				if isValidUrl(msg.Text) {
					newOrder.Mockup = msg.Text
					newOrder.state = stateDeadline
					text = steps[stateDeadline]
					_, err = bot.Send(msg.Sender, text)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	})
	bot.Handle(&finOrderBtn, func(msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		newOrder.state = 0
		newOrder.state = stateDone
		putOrder(*newOrder, db)
		text := steps[stateDone]
		_, err := bot.Send(msg.Sender, text)
		if err != nil {
			log.Fatal(err)
		}
	})

	bot.Handle(&BackBtn, func (msg *tb.Message){
		newOrder := getOrder(msg.ID, msg.Sender.Username)
		if newOrder.state == stateCols {
			if newOrder.Type == "Blank" {
				newOrder.state = stateProdtype
			} else {
				newOrder.state = stateFeature1
			}
		} else if newOrder.state == stateFeature1 {
			if newOrder.ProductName == "Hoodie" {
				newOrder.state = stateFeature2
			} else {
				newOrder.state = stateProdtype
			}
		} else {
			newOrder.state--
		}
	})


	bot.Handle(tb.OnPhoto, func (msg *tb.Message){
		newOrder := getOrder(msg.Sender.ID, msg.Sender.Username)
		if newOrder.state == stateMock{
			newOrder.Mockup = msg.Photo.FileID
			newOrder.state = stateDeadline
			text := steps[stateDeadline]
			_, err = bot.Send(msg.Sender, text)
			if err != nil {
				log.Println(err)
			}
		}
		if newOrder.state == stateLayout{
			newOrder.Layout = msg.Photo.FileID
			newOrder.state = stateMock
			text := steps[stateMock]
			_, err = bot.Send(msg.Sender, text)
			if err != nil {
				log.Println(err)
			}
		}
	})
	bot.Handle("/photo", func (msg *tb.Message){
		p := &tb.Photo{File: tb.FromURL("https://helpx.adobe.com/content/dam/help/en/photoshop-elements/using/optimizing-images-jpeg-format/_jcr_content/main-pars/image_0/we_02.png")}
		_, err = bot.Send(msg.Sender, p)
		if err != nil {
		    log.Println(err)
		}
	})
	bot.Start()
}
