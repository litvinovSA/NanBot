package main

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"strconv"
	"strings"
)

func printOrder(order *Order) string {
	var orderPrint string
	orderPrint += "Тип заказа: " + fmt.Sprintln(l10n[order.Type])
	orderPrint += "Изделие: " + fmt.Sprintln(l10n[order.ProductName])
	if len(order.Features) != 0 {
		orderPrint += "Особенности: "
		for _, feature := range order.Features {
			orderPrint += l10n[feature] + ", "
		}
		orderPrint += "\n"
	}
	orderPrint += "Количество: " + fmt.Sprintln(order.Amount)
	orderPrint += "Количество цветов: " + fmt.Sprintln(order.Cols)
	orderPrint += "Срок: " + fmt.Sprintln(order.Deadline)
	orderPrint += "Комментарий: " + fmt.Sprintln(order.Comment)
	return orderPrint
}

func Serve(update tgbotapi.Update, newOrder *Order, id int64) tgbotapi.MessageConfig {
	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
		switch update.CallbackQuery.Data {
		case "T-shirt", "Hoodie", "Sweatshirt":
			newOrder.ProductName = update.CallbackQuery.Data
			msg.Text = "Расскажи какой формат заказа тебе нужен"
			msg.ReplyMarkup = orderTypeKeyboard
			break
		case "Blank", "Sewing":
			newOrder.Type = update.CallbackQuery.Data
			if newOrder.Type == "Sewing" {
				msg.Text = "Уточним некоторые детали. "
				switch newOrder.ProductName {
				case "T-shirt":
					msg.Text += "Футболки какого размера нужны?"
					msg.ReplyMarkup = tshirt
				case "Hoodie":
					msg.Text = "Выбирай варианты:"
					msg.ReplyMarkup = hoodie
				case "Sweatshirt":
					msg.Text += "Что с рукавами?"
					msg.ReplyMarkup = sweatshirt
				}
			} else {
				newOrder.Cols = 1
				msg.Text = "Сколько нужно штук?"
			}
		case "hoodieReglan", "hoodieOversize", "hoodieDefault":
			newOrder.Features = append(newOrder.Features, update.CallbackQuery.Data[len("hoodie"):])
			msg.Text = "Еще немного"
			msg.ReplyMarkup = pocket
		case "Reglan", "Set-in", "Default", "Oversize", "pocketSewing", "pocketSet-in":
			newOrder.Features = append(newOrder.Features, update.CallbackQuery.Data)
			msg.Text = "Разберемся с количеством цветов. Сколько их будет? [1-8]"
		case "Edit":
			msg.Text = "Что будем менять?"
			//msg.ReplyMarkup = editFieldPicker
		case "editPrint":
			msg.Text = "Вот твой заказ: \n" + printOrder(newOrder)
			//msg.ReplyMarkup = editFieldPicker
		case "1":
			msg.Text = "Сколько штук?"
		case "2":
			msg.Text = "Сколько цветов?"
		case "3":
		case "4":
		case "5":
		case "6":
		case "Finish":
			msg.Text = "Заказ принят. \n Оформить еще один?"
			msg.ReplyMarkup = finishKeyboard
		case "End":
			msg.Text = "Спасибо за заказ!"
		}
		return msg
	}

	if update.Message != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		if update.Message.Photo != nil {
			if newOrder.Layout == "" {
				newOrder.Layout = (*update.Message.Photo)[0].FileID
				msg.Text = "А теперь мокап"
			} else if newOrder.Mockup == "" {
				newOrder.Mockup = (*update.Message.Photo)[0].FileID
				msg.Text = "Что по датам? Когда все должно быть готово?"
			} else {
				msg.Text = "Эй, хватит ломать меня! Что с датами?"
			}
			return msg
		}
		if update.Message.Text != "" {
			if strings.ToLower(update.Message.Text) == "new" {
				msg.Text = "Окей, давай определимся какой же тип изделия тебе нужен. Футболки, худи или свитшоты?"
				msg.ReplyMarkup = typeKeyboard
			} else if number, err := strconv.Atoi(update.Message.Text); err == nil {
				if (newOrder.Type == "Sewing" && newOrder.Cols == 0) || newOrder.edit {
					newOrder.Cols = number
					if newOrder.edit {
						//msg.ReplyMarkup = editFieldPicker
					} else {
						msg.Text = "Сколько штук нужно?"
					}
				} else if newOrder.Amount == 0 || newOrder.edit {
					newOrder.Amount = number
					if newOrder.edit {
						//msg.ReplyMarkup = editFieldPicker
					} else {
						msg.Text = "Теперь мне нужно фото макета"
					}
				}

			} else {
				if newOrder.Deadline == "" {
					newOrder.Deadline = update.Message.Text
					msg.Text = "Последний шаг! Есть какие-то особенные комментарии?"
				} else if newOrder.Comment == "" {
					newOrder.Comment = update.Message.Text
					msg.Text = "Готово. Давай проверим, что все верно.\n"
					msg.Text += printOrder(newOrder)
					msg.ReplyMarkup = editPicker
				}
			}
			return msg
		}
	}
	return tgbotapi.NewMessage(id, "Упс, что-то пошло не так! Попробуй еще раз.")
}

func NewServe(update tgbotapi.Update, newOrder *Order, id int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(id, "Упс, что-то пошло не так! Попробуй еще раз.")
	if update.Message != nil {
		if update.Message.Photo != nil {
			//TODO: get photos
		} else if update.Message.Text != "" {
			newOrder.state++
			switch update.Message.Text {
			case l10n["T-shirt"], l10n["Sweatshirt"], l10n["Hoodie"]:
				if newOrder.state == 1 {

				} else {
					msg.ReplyMarkup = typeKeyboard
					newOrder.state--
				}

			}
		}
	}
}