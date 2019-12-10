package main

import tgbotapi "github.com/Syfaro/telegram-bot-api"

var availableCommand = map[string]string{
	"/start": "Start the bot",
}

var orderTypeKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Пошив", "Sewing"),
		tgbotapi.NewInlineKeyboardButtonData("Бланк", "Blank"),
	),
)

var typeKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Футболка", "T-shirt"),
		tgbotapi.NewInlineKeyboardButtonData("Худи", "Hoodie"),
		tgbotapi.NewInlineKeyboardButtonData("Свитшот", "Sweatshirt"),
	),
)

var tshirt = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Обычная", "Default"),
		tgbotapi.NewInlineKeyboardButtonData("Оверсайз", "Oversize"),
	),
)

var sweatshirt = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Втачной рукав", "Set-in"),
		tgbotapi.NewInlineKeyboardButtonData("Реглан", "Reglan"),
	),
)

var hoodie = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Обычное", "hoodieDefault"),
		tgbotapi.NewInlineKeyboardButtonData("Реглан", "hoodieReglan"),
		tgbotapi.NewInlineKeyboardButtonData("Оверсайз", "hoodieOversize"),
	),
)

var pocket = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Нашив карманов", "Sewing"),
		tgbotapi.NewInlineKeyboardButtonData("Втачной карман", "Set-in"),
	),
)

var colorNumKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "1"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
		tgbotapi.NewInlineKeyboardButtonData("7", "7"),
		tgbotapi.NewInlineKeyboardButtonData("8", "8"),
	),
)

var editPicker = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Все верно", "Finish"),
		tgbotapi.NewInlineKeyboardButtonData("Хочу изменить", "Edit"),
	),
)

var editFieldPicker = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Количество", "1"),
		tgbotapi.NewInlineKeyboardButtonData("Количество цветов", "2"),
		tgbotapi.NewInlineKeyboardButtonData("Сроки", "5"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Макет", "3"),
		tgbotapi.NewInlineKeyboardButtonData("Мокап", "4"),
		tgbotapi.NewInlineKeyboardButtonData("Комментарий", "6"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Все верно", "Finish"),
	),
)
var finishKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Оформить еще один заказ", "New"),
		tgbotapi.NewInlineKeyboardButtonData("Закончить", "End"),
	),
)
