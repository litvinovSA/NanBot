package main

var l10n = map[string]string{
	"Sewing":         "Пошив",
	"Blank":          "Бланк",
	"T-shirt":        "Футболка",
	"Hoodie":         "Худи",
	"Sweatshirt":     "Свитшот",
	"Default":        "Обычная",
	"Set-in":         "Втачной рукав",
	"Reglan":         "Реглан",
	"pocketSewing":   "Нашив карманов",
	"pocketSet-in":   "Втачной карман",
	"Oversize":       "Оверсайз",
	"hoodieReglan":   "Реглан",
	"hoodieOversize": "Оверсайз",
	"hoodieDefault":  "Обычное",
	"adminNew":       "Новые",
	"adminProgress":  "В производстве",
	"adminDone":      "Сделанные",
	"Done":           "Закончить",
	"jpeg":           "Растровое изображение",
	"dunno":          "Я не знаю",
	"Back":           "Назад",
	"AmountBlank":    "Какое количество изделий? (Тираж)",
	"Thanks":         "Спасибо за заказ, он уже отправлен людям! Сегодня мы с вами свяжемся 🤗",
	"NewOrder":       "Оформить заказ",
	"Contacts":       "Контакты",
	"Orders":         "Мои заказы",
	"new":            "Новый",
	"inprogress":     "В производстве",
	"done":           "Выполнен",
	"another":        "Оформить еще один заказ",
}

var ru2eng = map[string]string{
	"Пошив":                   "Sewing",
	"Бланк":                   "Blank",
	"Футболка":                "T-shirt",
	"Худи":                    "Hoodie",
	"Свитшот":                 "Sweatshirt",
	"Обычная":                 "Default",
	"Втачной рукав":           "Set-in",
	"Реглан":                  "Reglan",
	"Нашив карманов":          "pocketSewing",
	"Втачной карман":          "pocketSet-in",
	"Оверсайз":                "Oversize",
	"Реглан худи":             "hoodieReglan",
	"Оверсайз худи":           "hoodieOversize",
	"Обычное худи":            "hoodieDefault",
	"Показать новые":          "adminNew",
	"Показать сделанные":      "adminDone",
	"Показать в производстве": "adminProgress",
}

var steps = map[int]string{
	stateHello:    "Привет, оформляем заказ!",
	stateProduct:  "Выберите тип изделия\n(Если заходите оформить еще один тип изделия, я спрошу вас позже)",
	stateProdtype: "Будем шить или выберем готовый бланк?",
	stateFeature1: "Уточним детали!",
	stateFeature2: "Какое именно будет худи?",
	stateCols:     "Какое количество цветов в принте? ",
	stateAmount:   "Сколько шьем?",
	stateLayout:   "Теперь добавьте мокап (как это должно выглядеть).\nИзображение или ссылка",
	stateMock:     "Нужен макет (что печатаем).\nИзображение или ссылка",
	stateDeadline: "Когда все должно быть готово? (Формат: дд/мм/гггг)",
	stateComment:  "Комментарии к заказу и особые пожелания? ☺️",
	stateFin:      "Вот ваш заказ\n",
	stateDone:     "Спасибо за заказ, он уже отправлен людям! Сегодня мы с вами свяжемся ",
}
