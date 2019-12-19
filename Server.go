package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/jmoiron/sqlx"
	"log"
	"net/url"
	"strconv"
)

const (
	stateHello    = iota
	stateProduct  = iota
	stateProdtype = iota
	stateFeature1 = iota
	stateFeature2 = iota
	stateCols     = iota
	stateAmount   = iota
	stateMock     = iota
	stateLayout   = iota
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

func adminServe(bot *tgbotapi.BotAPI, update tgbotapi.Update, id int64, db *sqlx.DB) tgbotapi.MessageConfig {
	if update.Message != nil {
		if update.Message.Text != "" {
			msg := tgbotapi.NewMessage(id, "")
			switch update.Message.Text {
			case ru2eng["adminNew"]:
				newOrders := getNewOrders(db)
				for _, order := range newOrders {
					msg.Text = stringifyOrder(&order)
					msg.ReplyMarkup = orderChangeStatus
					_, err := bot.Send(msg)
					if err != nil {
						log.Fatal(err)
					}
				}

			case ru2eng["adminProgress"]:
			case ru2eng["adminDone"]:
			}
		}
	}
	return tgbotapi.NewMessage(id, "Привет, Мастер!")
}

func NewServe(update tgbotapi.Update, newOrder *Order, id int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(id, "Упс, что-то пошло не так! Попробуй еще раз.")
	if update.Message != nil {
		if update.Message.Photo != nil {
			switch newOrder.state {
			case stateLayout:
				newOrder.Layout = (*update.Message.Photo)[0].FileID
				msg.Text = steps["Mock"]
				newOrder.state = stateMock
			case stateMock:
				newOrder.Mockup = (*update.Message.Photo)[0].FileID
				msg.Text = steps["Deadline"]
				newOrder.state = stateDeadline
			}
		} else if update.Message.Text != "" {
			switch update.Message.Text {
			case l10n["T-shirt"], l10n["Sweatshirt"], l10n["Hoodie"]:
				if newOrder.state == stateProduct {
					newOrder.ProductName = ru2eng[update.Message.Text]
					msg.Text = steps["ProdType"]
					msg.ReplyMarkup = orderTypeKeyboard
					newOrder.state = stateProdtype
				}
			case l10n["Blank"], l10n["Sewing"]:
				if newOrder.state == stateProdtype {
					newOrder.Type = ru2eng[update.Message.Text]
					if update.Message.Text == l10n["Blank"] {
						newOrder.state = stateCols
						msg.Text = steps["Cols"]
					} else {
						msg.Text = steps["Feature1"]
						switch newOrder.ProductName {
						case "T-shirt":
							newOrder.state = stateFeature1
							msg.Text = steps["Feature1"]
							msg.ReplyMarkup = tshirt
						case "Sweatshirt":
							newOrder.state = stateFeature1
							msg.Text = steps["Feature1"]
							msg.ReplyMarkup = sweatshirt
						case "Hoodie":
							newOrder.state = stateFeature2
							msg.Text = steps["Feature2"]
							msg.ReplyMarkup = hoodie
						}
					}
				}
			default:
				switch newOrder.state {
				case stateFeature1:
					if update.Message.Text == l10n["Default"] ||
						update.Message.Text == l10n["Oversize"] ||
						update.Message.Text == l10n["Set-in"] ||
						update.Message.Text == l10n["Reglan"] ||
						update.Message.Text == l10n["pocketSewing"] ||
						update.Message.Text == l10n["pocketSet-in"] {
						newOrder.Features += ru2eng[update.Message.Text]
						msg.Text = steps["Cols"]
						newOrder.state = stateCols
					}
				case stateFeature2:
					if update.Message.Text == l10n["hoodieDefault"] ||
						update.Message.Text == l10n["Reglan"] ||
						update.Message.Text == l10n["Oversize"] {
						newOrder.Features += ru2eng[update.Message.Text]
						msg.Text = steps["Feature1"]
						newOrder.state = stateFeature1
					}
				case stateCols:
					// here we can recieve photo or numbers
					if number, err := strconv.Atoi(update.Message.Text); err == nil {
						if number > 8 || number < 1 {
							msg.Text = "Хеллоу! От одного до восьми, ёпта!"
						} else {
							newOrder.Cols = number
							msg.Text = steps["Amount"]
							newOrder.state = stateAmount
						}
					}
				case stateAmount:
					if number, err := strconv.Atoi(update.Message.Text); err == nil {
						newOrder.Amount = number
						msg.Text = steps["Amount"]
						newOrder.state = stateLayout
					}
				case stateLayout:
					if isValidUrl(update.Message.Text) {
						newOrder.Layout = update.Message.Text
						msg.Text = steps["Mock"]
						newOrder.state = stateMock
					}
				case stateMock:
					if isValidUrl(update.Message.Text) {
						newOrder.Mockup = update.Message.Text
						msg.Text = steps["Deadline"]
						newOrder.state = stateDeadline
					}
				case stateDeadline:
					newOrder.Deadline = update.Message.Text
					msg.Text = steps["Comment"]
					newOrder.state = stateComment
				case stateComment:
					newOrder.Comment = update.Message.Text
					msg.Text = "Вот твой заказ: \n" + stringifyOrder(newOrder)
					newOrder.state = stateFin
					msg.ReplyMarkup = finishKeyboard
				}
			}
		}
	}
	return msg
}
