package main

import tgbotapi "github.com/Syfaro/telegram-bot-api"

var availableCommand = map[string]string{
	"/start": "Start the bot",
}

//var Keyboards = map[int]tgbotapi.ReplyKeyboardMarkup{
//	stateHello    : ,
//	stateProduct  : ,
//	stateProdtype : ,
//	stateFeature1 : ,
//	stateFeature2 : ,
//	stateCols     : ,
//	stateAmount   : ,
//	stateMock     : ,
//	stateLayout   : ,
//	stateDeadline : ,
//	stateComment  : ,
//	stateFin      : ,
//}

var orderTypeKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["Sewing"]),
		tgbotapi.NewKeyboardButton(l10n["Blank"]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["Back"]),
	),
)

var typeKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Футболка"),
		tgbotapi.NewKeyboardButton("Худи"),
		tgbotapi.NewKeyboardButton("Свитшот"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["Back"]),
	),
)

var tshirt = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Обычная"),
		tgbotapi.NewKeyboardButton("Оверсайз"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["Back"]),
	),
)

var sweatshirt = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Втачной рукав"),
		tgbotapi.NewKeyboardButton("Реглан"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["Back"]),
	),
)

var hoodie = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Обычное"),
		tgbotapi.NewKeyboardButton("Реглан"),
		tgbotapi.NewKeyboardButton("Оверсайз"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["Back"]),
	),
)
var adminStart = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Показать новые заказы"),
		tgbotapi.NewKeyboardButton("Показать в процессе"),
		tgbotapi.NewKeyboardButton("Показать сделанные"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["Back"]),
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
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["Back"]),
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
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["Back"]),
	),
)

var Cols = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
		tgbotapi.NewKeyboardButton("4"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
		tgbotapi.NewKeyboardButton("7"),
		tgbotapi.NewKeyboardButton("8"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["jpeg"]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["Back"]),
	),
)

var defaultKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(l10n["Back"]),
	),
)
