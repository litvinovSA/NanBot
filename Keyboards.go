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

var pocket = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Нашив карманов"),
		tgbotapi.NewKeyboardButton("Втачной карман"),
	),
)

//
//var colorNumKeyboard = tgbotapi.NewReplyKeyboard(
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("1", "1"),
//		tgbotapi.NewKeyboardButton("2", "2"),
//		tgbotapi.NewKeyboardButton("3", "3"),
//		tgbotapi.NewKeyboardButton("4", "4"),
//	),
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("5", "5"),
//		tgbotapi.NewKeyboardButton("6", "6"),
//		tgbotapi.NewKeyboardButton("7", "7"),
//		tgbotapi.NewKeyboardButton("8", "8"),
//	),
//)

var editPicker = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Все верно"),
		tgbotapi.NewKeyboardButton("Хочу изменить"),
	),
)

//
//var editFieldPicker = tgbotapi.NewReplyKeyboard(
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("Количество", "1"),
//		tgbotapi.NewKeyboardButton("Количество цветов", "2"),
//tgbotapi.NewKeyboardButton("Сроки", "5"),
//),
//tgbotapi.NewKeyboardButtonRow(
//	tgbotapi.NewKeyboardButton("Макет", "3"),
//	tgbotapi.NewKeyboardButton("Мокап", "4"),
//	tgbotapi.NewKeyboardButton("Комментарий", "6"),
//),
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("Показать заказ", "editPrint"),
//	),
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("Все верно", "Finish"),
//	),
//)
var finishKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Оформить еще один заказ"),
		tgbotapi.NewKeyboardButton("Закончить"),
	),
)
