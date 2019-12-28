package main

import (
	"log"
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)


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
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
            ReplyKeyboard:startKeyboard,
            ResizeReplyKeyboard:true,
        })
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
			file, err := bot.FileByID(order.Layout)
			if err != nil {
				log.Println(err)
			}
			file1, err := bot.FileByID(order.Mockup)
			
			if err != nil {
			    log.Println(err)
			}
			album := tb.Album{&tb.Photo{File: file}, &tb.Photo{File: file1}}
			_, err = bot.SendAlbum(msg.Sender, album)
			
			if err != nil {
			    log.Println(err)
			}
			
			text := stringifyOrder(&order) + "\n Статус: " + l10n[order.State]
			_, err = bot.Send(msg.Sender, text)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
	bot.Handle(&newOrderBtn, func(msg *tb.Message){
		orders[msg.Sender.ID] = initOrder(msg.Sender.Username, msg.Sender.ID, getNextID(db))
		text := steps[stateType]
		orders[msg.Sender.ID].state = stateType
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ReplyKeyboard:orderTypeKeyboard,
			ResizeReplyKeyboard:true})
		if err != nil {
			log.Fatal(err)
		}
	})
	// Block 2 order type
	bot.Handle(&sewingBtn, func(msg *tb.Message){
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
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
	// Block 3 product
	bot.Handle(&TshirtBtn, func(msg *tb.Message){
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
		newOrder.state = 0
		newOrder.state = stateAmount
		text := steps[stateAmount]
		_, err := bot.Send(msg.Sender, text)
		if err != nil {
			log.Fatal(err)
		}
	})
	bot.Handle(&jpegBtn, func(msg *tb.Message){
		newOrder := orders[msg.Sender.ID]
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
		newOrder := orders[msg.Sender.ID]
		keys := tb.ReplyMarkup{
			ReplyKeyboard:defaultKeyboard,
			ResizeReplyKeyboard:true,
			OneTimeKeyboard:true,
		}
		var text string
		if num, err := strconv.Atoi(msg.Text); err == nil {
			if newOrder.state == stateCols {
				newOrder.Cols = num
				newOrder.state = stateAmount
				if newOrder.Type == "Blank" {
					text = l10n["AmountBlank"]
				} else {
					text = steps[stateAmount]
				}
				_, err = bot.Send(msg.Sender, text, &keys)
				if err != nil {
				    log.Println(err)
				}
			} else if newOrder.state == stateAmount {
				newOrder.Amount = num
				text = steps[stateLayout]
				newOrder.state = stateLayout
				_, err = bot.Send(msg.Sender, text, &keys)
				if err != nil {
				    log.Println(err)
				}
			}
		} else {
			if newOrder.state == stateDeadline {
				newOrder.Deadline = msg.Text
				newOrder.state = stateComment
				text = steps[stateComment]
				_, err = bot.Send(msg.Sender, text, &keys)
				if err != nil {
				    log.Println(err)
				}
			} else if newOrder.state == stateComment{
				newOrder.Comment = msg.Text
				newOrder.state = stateFin
				text = steps[stateFin]
				text += "\n" + stringifyOrder(newOrder)
				if len(newOrder.molAlbum) != 0 {
					_, err = bot.SendAlbum(msg.Sender, newOrder.molAlbum)
					if err != nil {
						log.Println(err)
					}
				}
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
		newOrder := orders[msg.Sender.ID]
		newOrder.state = 0
		newOrder.state = stateDone
		go putOrder(*newOrder, db)
		text := steps[stateDone]
		_, err := bot.Send(msg.Sender, text, &tb.ReplyMarkup{
			ResizeReplyKeyboard:true,
			ReplyKeyboard: startKeyboard,
		})
		if err != nil {
			log.Fatal(err)
		}
	})

	bot.Handle(&BackBtn, func (msg *tb.Message){
		newOrder := orders[msg.Sender.ID]
		if newOrder.state == stateCols {
			if newOrder.Type == l10n["Blank"] {
				newOrder.state = stateProduct
			} else {
				newOrder.state = stateFeature1
			}
		} else if newOrder.state == stateFeature1 {
			if newOrder.ProductName == l10n["Hoodie"] {
				newOrder.state = stateFeature2
			} else {
				newOrder.state = stateProduct
			}
		} else {
			newOrder.state--
		}
		text, keys := getKeyboardAndTextByState(newOrder)
		_, err := bot.Send(msg.Sender, text, &keys)
		if err != nil {
		    log.Println(err)
		}

	})

	bot.Handle(tb.OnPhoto, func (msg *tb.Message){
		newOrder := orders[msg.Sender.ID]
		if newOrder.state == stateMock{
			newOrder.Mockup = msg.Photo.FileID
			newOrder.molAlbum = append(newOrder.molAlbum, msg.Photo)
			newOrder.state = stateDeadline
			text := steps[stateDeadline]
			_, err = bot.Send(msg.Sender, text)
			if err != nil {
				log.Println(err)
			}
		}
		if newOrder.state == stateLayout{
			newOrder.Layout = msg.Photo.FileID
			newOrder.molAlbum = append(newOrder.molAlbum, msg.Photo)
			newOrder.state = stateMock
			text := steps[stateMock]
			_, err = bot.Send(msg.Sender, text)
			if err != nil {
				log.Println(err)
			}
		}
	})
	bot.Handle(&anotherBtn, func (msg *tb.Message){
		newOrder := orders[msg.Sender.ID]
		go putOrder(*newOrder, db)
		newOrder.state = stateType
		_, err = bot.Send(msg.Sender, steps[newOrder.state], &tb.ReplyMarkup{
			ReplyKeyboard: typeKeyboard,
			ResizeReplyKeyboard: true,
			OneTimeKeyboard: true,
		})
		if err != nil {
		    log.Println(err)
		}
	})
	bot.Handle("/admin", func(msg *tb.Message){
		if msg.Sender.Username == "mrdken" || msg.Sender.Username == "potishebud"{
			bot.Send(msg.Sender, l10n["admin"], &tb.ReplyMarkup{
				ReplyKeyboard:       adminKeybord,
				ResizeReplyKeyboard: true,
			})
		} else {
			bot.Send(msg.Sender, l10n["not_admin"], &tb.ReplyMarkup{
				ReplyKeyboard:       startKeyboard,
				ResizeReplyKeyboard: true,
			})
		}
	})
	bot.Handle(&adminNew, func(msg *tb.Message){
		orders := getNewOrders(db)
		for _, order := range orders {
			file, err := bot.FileByID(order.Layout)
			if err != nil {
				log.Println(err)
			}
			file1, err := bot.FileByID(order.Mockup)

			if err != nil {
				log.Println(err)
			}
			album := tb.Album{&tb.Photo{File: file}, &tb.Photo{File: file1}}
			_, err = bot.SendAlbum(msg.Sender, album)

			if err != nil {
				log.Println(err)
			}

			text := tb.Message{Text:stringifyOrder(&order) + "\n Статус: " + l10n[order.State],
				ReplyMarkup: tb.InlineKeyboardMarkup{InlineKeyboard:adminInlineKeyboard}}
			_, err = bot.Send(msg.Sender, text)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
	bot.Handle(&adminProgress, func(msg *tb.Message){
		orders := getInProgresOrders(db)
		for _, order := range orders {
			file, err := bot.FileByID(order.Layout)
			if err != nil {
				log.Println(err)
			}
			file1, err := bot.FileByID(order.Mockup)

			if err != nil {
				log.Println(err)
			}
			album := tb.Album{&tb.Photo{File: file}, &tb.Photo{File: file1}}
			_, err = bot.SendAlbum(msg.Sender, album)

			if err != nil {
				log.Println(err)
			}

			text := stringifyOrder(&order) + "\n Статус: " + l10n[order.State]
			_, err = bot.Send(msg.Sender, text)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
	bot.Handle(&adminDone, func(msg *tb.Message){
		orders := getDoneOrders(db)
		for _, order := range orders {
			file, err := bot.FileByID(order.Layout)
			if err != nil {
				log.Println(err)
			}
			file1, err := bot.FileByID(order.Mockup)

			if err != nil {
				log.Println(err)
			}
			album := tb.Album{&tb.Photo{File: file}, &tb.Photo{File: file1}}
			_, err = bot.SendAlbum(msg.Sender, album)

			if err != nil {
				log.Println(err)
			}

			text := stringifyOrder(&order) + "\n Статус: " + l10n[order.State]
			_, err = bot.Send(msg.Sender, text)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
	bot.Handle(&adminToDone, func(msg *tb.Message){
		id := parseId(msg.Text)
		moveToDone(id, db)
		err := bot.Delete(msg)
		if err != nil {
		    log.Println(err)
		}
	})
	bot.Handle(&adminToDone, func(msg *tb.Message){
		id := parseId(msg.Text)
		moveToDone(id, db)
		err := bot.Delete(msg)
		if err != nil {
			log.Println(err)
		}
	})


	bot.Handle("/photo", func (msg *tb.Message){
		file, err := bot.FileByID("AgADAgADhKwxG3aMOEh1WwzSyDzx4KbEwg8ABAEAAwIAA3kAA750AQABFgQ")
		if err != nil {
		    log.Println(err)
		}
		photo := &tb.Photo{File: file}
		bot.Send(msg.Sender, photo)
	})
	bot.Start()
}
