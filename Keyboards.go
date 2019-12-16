package main

import tgbotapi "github.com/Syfaro/telegram-bot-api"

var availableCommand = map[string]string{
	"/start": "Start the bot",
}

var orderTypeKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Пошив", "Sewing"),
		tgbotapi.NewKeyboardButton("Бланк", "Blank"),
	),
)

var typeKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Футболка", "T-shirt"),
		tgbotapi.NewKeyboardButton("Худи", "Hoodie"),
		tgbotapi.NewKeyboardButton("Свитшот", "Sweatshirt"),
	),
)

var tshirt = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Обычная", "Default"),
		tgbotapi.NewKeyboardButton("Оверсайз", "Oversize"),
	),
)

var sweatshirt = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Втачной рукав", "Set-in"),
		tgbotapi.NewKeyboardButton("Реглан", "Reglan"),
	),
)

var hoodie = tgbotapi.keyboardNewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Обычное", "hoodieDefault"),
		tgbotapi.NewKeyboardButton("Реглан", "hoodieReglan"),
		tgbotapi.NewKeyboardButton("Оверсайз", "hoodieOversize"),
	),
)

var pocket = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Нашив карманов", "pocketSewing"),
		tgbotapi.NewKeyboardButton("Втачной карман", "pocketSet-in"),
	),
)

var colorNumKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1", "1"),
		tgbotapi.NewKeyboardButton("2", "2"),
		tgbotapi.NewKeyboardButton("3", "3"),
		tgbotapi.NewKeyboardButton("4", "4"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("5", "5"),
		tgbotapi.NewKeyboardButton("6", "6"),
		tgbotapi.NewKeyboardButton("7", "7"),
		tgbotapi.NewKeyboardButton("8", "8"),
	),
)

var editPicker = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Все верно", "Finish"),
		tgbotapi.NewKeyboardButton("Хочу изменить", "Edit"),
	),
)

var editFieldPicker = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Количество", "1"),
		tgbotapi.NewKeyboardButton("Количество цветов", "2"),
		//tgbotapi.NewKeyboardButton("Сроки", "5"),
	),
	//tgbotapi.NewKeyboardButtonRow(
	//	tgbotapi.NewKeyboardButton("Макет", "3"),
	//	tgbotapi.NewKeyboardButton("Мокап", "4"),
	//	tgbotapi.NewKeyboardButton("Комментарий", "6"),
	//),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Показать заказ", "editPrint"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Все верно", "Finish"),
	),
)
var finishKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Оформить еще один заказ", "New"),
		tgbotapi.NewKeyboardButton("Закончить", "End"),
	),
)
