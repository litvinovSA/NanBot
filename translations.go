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

var ru2eng = map [string]string{
	"Пошив": "Sewing",
	"Бланковое изделие": "Blank",
	"Футболка":        "T-shirt",
	"Худи":         "Hoodie",
	"Свитшот""Sweatshirt":     "Свитшот",
	"Обычная""Default":        ,
	"Втачной рукав""Set-in":         ,
	"Реглан""Reglan":         ,
	"pocketSewing":   "Нашив карманов",
	"pocketSet-in":   "Втачной карман",
	"Oversize":       "Оверсайз",
	"hoodieReglan":   "Реглан",
	"hoodieOversize": "Оверсайз",
	"hoodieDefault":  "Обычное",
}

var steps = map[string]string{
	"Greeting" : "Привет, оформляем заказ!",
	"Product"  : "Выбери тип изделия",
	"ProdType" : "Расскажи, какой формат заказа тебе нужен?",
	"Feature1" : "Какие будет особенности?",
	"Feature2" : "",
	"Cols" : "" ,
	"Amount" : "",
	"Mock": "",
	"Layout": "",
	"Deadline": "",
	"Comment" : "",
}