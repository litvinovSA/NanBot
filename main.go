package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"strconv"
)

type Order struct {
	id           int
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
}

func printOrder(order Order) string {
	return ""
}

func main() {
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
		if update.CallbackQuery != nil {
			fmt.Print(update)

			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data))
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			switch update.CallbackQuery.Data {
			case "T-shirt", "Hoodie", "Sweatshirt":
				newOrder.productName = update.CallbackQuery.Data
				msg.Text = "Расскажи какой формат заказа тебе нужен"
				msg.ReplyMarkup = orderTypeKeyboard
				break
			case "Blank", "Sewing":
				newOrder.orderType = update.CallbackQuery.Data
				if newOrder.orderType == "Sewing" {
					msg.Text = "Уточним некоторые детали. "
					switch newOrder.productName {
					case "T-shirt":
						msg.Text += "Футболки какого размера нужны?"
						msg.ReplyMarkup = tshirt
					case "Hoodie":
						msg.Text = "Пока не поддерживаем, попробуй другие варинаты"
						msg.ReplyMarkup = typeKeyboard
					case "Swearshirt":
						msg.Text += "Что с рукавами?"
						msg.ReplyMarkup = sweatshirt
					}
				}
			case "Reglan", "Set-in", "Default", "Oversize", "pocketSewing":
				newOrder.features = append(newOrder.features, update.CallbackQuery.Data)
				msg.Text = "Разберемся с количеством цветов. Сколько из будет? [1-8]"
			case "Edit":
				msg.Text = "Что будем менять?"
				msg.ReplyMarkup = editFieldPicker
			}
			bot.Send(msg)

		}

		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			if update.Message.Photo != nil {
				if newOrder.layout != "" {
					newOrder.layout = (*update.Message.Photo)[0].FileID
					msg.Text = "А теперь мокап"
				} else if newOrder.mockup != "" {
					newOrder.layout = (*update.Message.Photo)[0].FileID
					msg.Text = "Что по датам? Когда все должно быть готово?"
				}
			}
			switch update.Message.Text {
			case "new":
				newOrder.id = 1
				msg.Text = "Окей, давай определимся какой же тип изделия тебе нужен. Футболки, худи или свитшоты?"
				msg.ReplyMarkup = typeKeyboard
			default:
				if number, err := strconv.Atoi(update.Message.Text); err == nil {
					if newOrder.orderType == "Sewing" && newOrder.amountCol == 0 {
						newOrder.amountCol = number
					} else if newOrder.amount == 0 {
						newOrder.amount = number
						msg.Text = "Теперь мне нужно фото макета"
					}

				} else {
					if newOrder.deadline == "" {
						newOrder.deadline = update.Message.Text
						msg.Text = "Последний шаг! Есть какие-то особенные комментарии?"
					} else if newOrder.comment == "" {
						newOrder.deadline = update.Message.Text
						msg.Text = "Готово. Давай проверим, что все верно.\n"
						msg.Text += printOrder(newOrder)
						msg.ReplyMarkup = editPicker
					}
				}

			}

			bot.Send(msg)
		}
	}
}
