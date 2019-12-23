package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/jmoiron/sqlx"
	"log"
	"net/url"
	"strconv"
	"strings"
)

const (
	stateHello    = iota
	stateProduct  = iota
	stateProdtype = iota
	stateFeature1 = iota
	stateFeature2 = iota
	stateCols     = iota
	stateAmount   = iota
	stateLayout   = iota
	stateMock     = iota
	stateDeadline = iota
	stateComment  = iota
	stateFin      = iota
)

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}

func parseId(order string) int {
	id, err := strconv.Atoi(strings.Split(strings.Split(order, "\n")[0], ": ")[1])
	if err != nil {
		log.Print(err)
	}
	return id
}

func getPhoto(photoId string) {

}

func adminServe(bot *tgbotapi.BotAPI, update tgbotapi.Update, id int64, db *sqlx.DB) tgbotapi.MessageConfig {
	if update.CallbackQuery != nil {
		switch update.CallbackQuery.Data {
		case "Done":
			moveToDone(parseId(update.CallbackQuery.Message.Text), db)
		case "Production":
			moveToProgress(parseId(update.CallbackQuery.Message.Text), db)
		}
	} else if update.Message != nil {

		if update.Message.Text != "" {
			msg := tgbotapi.NewMessage(id, "")
			switch update.Message.Text {
			case "admin", "Admin":
				msg := tgbotapi.NewMessage(id, "Привет, Мастер!")
				msg.ReplyMarkup = adminDefault
				bot.Send(msg)
			case l10n["adminNew"]:
				newOrders := getNewOrders(db)
				if len(newOrders) == 0 {
					msg.Text = "Новых заказов нет. Вы молодец, Senpai! ^-^"
					bot.Send(msg)
				} else {
					for _, order := range newOrders {
						msg.Text = stringifyOrder(&order)
						msg.ReplyMarkup = orderChangeStatus
						_, err := bot.Send(msg)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			case l10n["adminProgress"]:
				orders := getInProgresOrders(db)
				for _, order := range orders {
					msg.Text = stringifyOrder(&order)
					_, err := bot.Send(msg)
					if err != nil {
						log.Fatal(err)
					}
				}
			case l10n["adminDone"]:
				orders := getDoneOrders(db)
				for _, order := range orders {
					msg.Text = stringifyOrder(&order)
					_, err := bot.Send(msg)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}

	return tgbotapi.NewMessage(id, "Привет, Мастер!")
}

func getKeyboardAndTextByState(newOrder *Order, id int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(id, l10n["Error"])
	msg.Text = steps[newOrder.state]
	switch newOrder.state {
	case stateProduct:
		msg.ReplyMarkup = typeKeyboard
	case stateProdtype:
		msg.ReplyMarkup = orderTypeKeyboard
	case stateFeature1:
		switch newOrder.ProductName {
		case "Hoodie":
			msg.ReplyMarkup = hoodie
		case "T-shirt":
			msg.ReplyMarkup = tshirt
		case "Sweatshirt":
			msg.ReplyMarkup = sweatshirt
		}
	case stateFeature2:
		msg.ReplyMarkup = pocket
	case stateCols:
		msg.ReplyMarkup = Cols
	case stateAmount:
		if newOrder.Type == "Blank" {
			msg.Text = l10n["AmountBlank"]
		}
		msg.ReplyMarkup = defaultKeyboard
	case stateLayout:
		msg.ReplyMarkup = defaultKeyboard
	case stateMock:
		msg.ReplyMarkup = defaultKeyboard
	case stateDeadline:
		msg.ReplyMarkup = defaultKeyboard
	case stateComment:
		msg.ReplyMarkup = defaultKeyboard
	case stateFin:
		msg.ReplyMarkup = finishKeyboard
	}
	return msg
}

func NewServe(update tgbotapi.Update, newOrder *Order, id int64, db *sqlx.DB) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(id, "Упс, что-то пошло не так! Попробуй еще раз.")
	if update.Message != nil {
		if update.Message.Photo != nil {
			switch newOrder.state {
			case stateLayout:
				newOrder.Layout = (*update.Message.Photo)[0].FileID
				newOrder.state = stateMock
				msg = getKeyboardAndTextByState(newOrder, id)
			case stateMock:
				newOrder.Mockup = (*update.Message.Photo)[0].FileID
				newOrder.state = stateDeadline
				msg = getKeyboardAndTextByState(newOrder, id)
			}
		} else if update.Message.Text != "" {
			switch update.Message.Text {
			case l10n["Back"]:
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
			case l10n["T-shirt"], l10n["Sweatshirt"], l10n["Hoodie"]:
				if newOrder.state == stateProduct {
					newOrder.ProductName = ru2eng[update.Message.Text]
					newOrder.state = stateProdtype
				}
			case l10n["Blank"], l10n["Sewing"]:
				if newOrder.state == stateProdtype {
					newOrder.Type = ru2eng[update.Message.Text]
					if update.Message.Text == l10n["Blank"] {
						newOrder.state = stateCols
					} else {
						switch newOrder.ProductName {
						case "T-shirt":
							newOrder.state = stateFeature1
						case "Sweatshirt":
							newOrder.state = stateFeature1
						case "Hoodie":
							newOrder.state = stateFeature2
						}
					}
				}
			case l10n["Done"]:
				go putOrder(*newOrder, db)
				msg.Text = l10n["Thanks"]
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			default:
				switch newOrder.state {
				case stateHello:
					newOrder.state = stateProduct
				case stateFeature1:
					if update.Message.Text == l10n["Default"] ||
						update.Message.Text == l10n["Oversize"] ||
						update.Message.Text == l10n["Set-in"] ||
						update.Message.Text == l10n["Reglan"] ||
						update.Message.Text == l10n["pocketSewing"] ||
						update.Message.Text == l10n["pocketSet-in"] {
						newOrder.Features += update.Message.Text
						newOrder.state = stateCols
					}
				case stateFeature2:
					if update.Message.Text == l10n["hoodieDefault"] ||
						update.Message.Text == l10n["Reglan"] ||
						update.Message.Text == l10n["Oversize"] {
						newOrder.Features += update.Message.Text + ", "
						newOrder.state = stateFeature1
					}
				case stateCols:
					if number, err := strconv.Atoi(update.Message.Text); err == nil {
						if update.Message.Text == l10n["jpeg"] {
							newOrder.Cols = -1
						} else {
							newOrder.Cols = number
						}
						newOrder.state = stateAmount
					}
				case stateAmount:
					if number, err := strconv.Atoi(update.Message.Text); err == nil {
						newOrder.Amount = number
						newOrder.state = stateLayout
					}
				case stateLayout:
					if isValidUrl(update.Message.Text) {
						newOrder.Layout = update.Message.Text
						newOrder.state = stateMock
					}
				case stateMock:
					if isValidUrl(update.Message.Text) {
						newOrder.Mockup = update.Message.Text
						newOrder.state = stateDeadline
					}
				case stateDeadline:
					newOrder.Deadline = update.Message.Text
					newOrder.state = stateComment
				case stateComment:
					newOrder.Comment = update.Message.Text
					newOrder.state = stateFin
				}
			}
		}
	}
	msg = getKeyboardAndTextByState(newOrder, id)
	return msg
}
