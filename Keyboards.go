package main

import tgbotapi "github.com/Syfaro/telegram-bot-api"

var availableCommand = map[string]string{
	"/start": "Start the bot",
}

var orderTypeKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Пошив"),
		tgbotapi.NewKeyboardButton("Бланк"),
	),
)

var typeKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Футболка"),
		tgbotapi.NewKeyboardButton("Худи"),
		tgbotapi.NewKeyboardButton("Свитшот"),
	),
)

var tshirt = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Обычная"),
		tgbotapi.NewKeyboardButton("Оверсайз"),
	),
)

var sweatshirt = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Втачной рукав"),
		tgbotapi.NewKeyboardButton("Реглан"),
	),
)

var hoodie = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Обычное"),
		tgbotapi.NewKeyboardButton("Реглан"),
		tgbotapi.NewKeyboardButton("Оверсайз"),
	),
)
var adminStart = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Показать новые заказы"),
		tgbotapi.NewKeyboardButton("Показать в процессе"),
		tgbotapi.NewKeyboardButton("Показать сделанные"),
	),
)

var orderChangeStatus = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("В производство", "Production"),
		tgbotapi.NewInlineKeyboardButtonData("Выполнен", "Done"),
	),
)

var pocket = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Нашив карманов"),
		tgbotapi.NewKeyboardButton("Втачной карман"),
	),
)

var adminDefault = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Новые"),
		tgbotapi.NewKeyboardButton("В производстве"),
		tgbotapi.NewKeyboardButton("Сделанные"),
	),
)

var editPicker = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Все верно"),
		tgbotapi.NewKeyboardButton("Хочу изменить"),
	),
)

var finishKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Оформить еще один заказ"),
		tgbotapi.NewKeyboardButton("Закончить"),
	),
)

var Cols = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["jpeg"]),
	),
)
