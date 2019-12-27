package main


import
(
	tb "gopkg.in/tucnak/telebot.v2"
)

var sewingBtn, blankBtn =
	tb.ReplyButton{Text:l10n["Sewing"]},
	tb.ReplyButton{Text:l10n["Blank"]}

var orderTypeKeyboard = [][]tb.ReplyButton{
	{sewingBtn, blankBtn},
	{BackBtn},
}

var BackBtn = tb.ReplyButton{Text: "Назад"}

var TshirtBtn = tb.ReplyButton{Text:"Футболка"}
var HoodBtn = tb.ReplyButton{Text:"Худи"}
var SweetBtn = tb.ReplyButton{Text:"Свитшот"}
var typeKeyboard = [][]tb.ReplyButton{
	{TshirtBtn, HoodBtn, SweetBtn},
}


var defaultTShirtBtn = tb.ReplyButton{Text:"Обычная"}
var oversizeTShirtBtn = tb.ReplyButton{Text:"Оверсайз"}
var tshirtKeyboard = [][]tb.ReplyButton{
	{defaultTShirtBtn,oversizeTShirtBtn},
	{BackBtn},
}

var SetinBtn = tb.ReplyButton{Text:"Втачной рукав"}
var sweetRelganBtn = tb.ReplyButton{Text:"Реглан"}
var sweetKeybord = [][]tb.ReplyButton{
	{SetinBtn,sweetRelganBtn},
	{BackBtn},
}
var hoodDefault = tb.ReplyButton{Text:"Обычное"}
var hoodReglan = tb.ReplyButton{Text:"Реглан"}
var hoodOversize = 	tb.ReplyButton{Text:"Оверсайз"}
var hoodKeyboard = [][]tb.ReplyButton{
	{hoodDefault,hoodReglan, hoodOversize},
	{BackBtn},
}

var pocketSewingBtn, pocketSetinBtn = tb.ReplyButton{Text:"Нашив карманов"},
	tb.ReplyButton{Text:"Втачной карман"}

var pocketKeyboard = [][]tb.ReplyButton{
	{pocketSewingBtn,pocketSetinBtn},
	{BackBtn},
}
var adminNew, adminProgress, adminDone =
	tb.ReplyButton{Text:"Новые"},
	tb.ReplyButton{Text:"В производстве"},
	tb.ReplyButton{Text:"Сделанные"}

var adminKeybord = [][]tb.ReplyButton{
	{adminNew,adminProgress, adminDone},

}
//
//var orderChangeStatus = tgbotapi.NewInlineKeyboardMarkup{
//	tgbotapi.NewInlineKeyboardRow{
//		tgbotapi.NewInlineKeyboardButtonData{"В производство", "Production"},
//		tgbotapi.NewInlineKeyboardButtonData{"Выполнен", "Done"},
//	},
//}
//
//var editPicker = tgbotapi.NewReplyKeyboard{
//	tgbotapi.NewKeyboardButtonRow{
//		tb.ReplyButton{Text:"Все верно"},
//		tb.ReplyButton{Text:"Хочу изменить"},
//	},
//}

var anotherBtn, finOrderBtn =
	tb.ReplyButton{Text:"Оформить еще один заказ"},
	tb.ReplyButton{Text:"Закончить"}

var finKeyboard = [][]tb.ReplyButton{
	[]tb.ReplyButton{anotherBtn,finOrderBtn},
	[]tb.ReplyButton{BackBtn},
}
var jpegBtn, dunnoBtn =
	tb.ReplyButton{Text:l10n["jpeg"]},
	tb.ReplyButton{Text:l10n["dunno"]}

var colorsKeyboard = [][]tb.ReplyButton{
	{	tb.ReplyButton{Text:"1"},
		tb.ReplyButton{Text:"2"},
		tb.ReplyButton{Text:"3"},
		tb.ReplyButton{Text:"4"},
	},
	{	tb.ReplyButton{Text:"5"},
		tb.ReplyButton{Text:"6"},
		tb.ReplyButton{Text:"7"},
		tb.ReplyButton{Text:"8"},
	},
	{jpegBtn},
	{dunnoBtn},
	{BackBtn},
}

var defaultKeyboard = [][]tb.ReplyButton{
	{BackBtn},
}

var newOrderBtn, contactsBtn, myordersBtn =
	tb.ReplyButton{Text:l10n["NewOrder"]},
	tb.ReplyButton{Text:l10n["Contacts"]},
	tb.ReplyButton{Text:l10n["Orders"]}

	var startKeyboard = [][]tb.ReplyButton{
	{newOrderBtn},
	{myordersBtn,contactsBtn},
}

