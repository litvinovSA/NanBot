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
	"jpeg":           "Я не знаю",
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

var steps = map[string]string{
	"Greeting":    "Привет, оформляем заказ!",
	"Product":     "Выберите тип изделия",
	"ProdType":    "Будем шить или выберем готовый бланк?",
	"Feature1":    "Уточним детали!",
	"Feature2":    "Еще немного деталей",
	"Cols":        "Какое количество цветов в принте? ",
	"Amount":      "Сколько шьем?",
	"AmountBlank": "Какое количество изделий? (Тираж)",
	"Mock":        "Теперь добавьте мокапом (как это должно выглядеть).\nИзображение или ссылка",
	"Layout":      "Нужен макет (что печатаем).\nИзображение или ссылка",
	"Deadline":    "Когда все должно быть готово? (Формат: дд/мм/гггг)",
	"Comment":     "Комментарии к заказу и особые пожелания? ☺️",
	"OrderString": "Вот ваш заказ\n",
	"Thanks":      "Спасибо за заказ, он уже отправлен людям! Сегодня мы с вами свяжемся 🤗",
}
