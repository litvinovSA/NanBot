package main

var l10n = map[string]string{
	"Sewing":         "Пошив",
	"Blank":          "Бланковое изделие",
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
}

var ru2eng = map[string]string{
	"Пошив":                   "Sewing",
	"Бланковое изделие":       "Blank",
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
	"Greeting": "Привет, оформляем заказ!",
	"Product":  "Выбери тип изделия",
	"ProdType": "Расскажи, какой формат заказа тебе нужен?",
	"Feature1": "Уточним детали!",
	"Feature2": "Еще немного деталей",
	"Cols":     "Количество цветов [1-8], либо растровое изображени",
	"Amount":   "Сколько шьем?",
	"Mock":     "Теперь дело за мокапом.\n Изображение или ссылка",
	"Layout":   "Нужен макет.\n Изображение или ссылка",
	"Deadline": "Когда все должно быть готово?",
	"Comment":  "Последние слова?",
}
