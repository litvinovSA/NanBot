package main

import tgbotapi "github.com/Syfaro/telegram-bot-api"

var availableCommand = map[string]string{
	"/start": "Start the bot",
}

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Пошив","Пошив"),
		tgbotapi.NewInlineKeyboardButtonData("Бланк","blank"),
	),
)

var typeKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Футболка","Пошив"),
		tgbotapi.NewInlineKeyboardButtonData("Худи","blank"),
		tgbotapi.NewInlineKeyboardButtonData("Свитшот","blank"),
	),
)

var tshirt = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Обычная","Пошив"),
		tgbotapi.NewInlineKeyboardButtonData("Оверсайз","blank"),
	),
)

var switshot = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Вточной рукав","Пошив"),
		tgbotapi.NewInlineKeyboardButtonData("Реглан","blank"),
	),
)

var hoody = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Обычное","Пошив"),
		tgbotapi.NewInlineKeyboardButtonData("Реглан","blank"),
		tgbotapi.NewInlineKeyboardButtonData("Оверсайз","blank"),
	),
)

var pocket = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Нашив карманов","Пошив"),
		tgbotapi.NewInlineKeyboardButtonData("Вточной карман","blank"),
	),
)

