package main

import (
	"fmt"
	"github.com/google/uuid"
)

type Order struct {
	Id          int
	Orderid     uuid.UUID `db:"orderid"`
	Type        string    `db:"type"`
	ProductName string    `db:"productname"`
	Features    string    `db:"features"`
	Amount      int       `db:"amount"`
	Cols        int       `db:"cols"`
	Mockup      string    `db:"mockup"`
	Layout      string    `db:"layout"`
	CustomerID  string    `db:"customerid"`
	Deadline    string    `db:"deadline"`
	Comment     string    `db:"comment"`
	State       string    `db:"state"`
	edit        bool
	state       int
}

var orders = make(map[int64]*Order)

func initOrder(username string, nextid int) *Order {
	return &Order{
		Id:          nextid,
		Orderid:     uuid.New(),
		Type:        "",
		ProductName: "",
		Features:    "",
		Amount:      0,
		Cols:        0,
		Mockup:      "",
		Layout:      "",
		CustomerID:  username,
		Deadline:    "",
		Comment:     "",
		State:       "new",
		edit:        false,
		state:       stateHello,
	}
}

func stringifyOrder(order *Order) string {
	var orderPrint string
	orderPrint += "Номер заказа: " + fmt.Sprintln(order.Id)
	orderPrint += "Ник заказчика: @" + fmt.Sprintln(order.CustomerID)
	orderPrint += "Тип заказа: " + fmt.Sprintln(l10n[order.Type])
	orderPrint += "Изделие: " + fmt.Sprintln(l10n[order.ProductName])
	if len(order.Features) != 0 {
		orderPrint += "Особенности: " + order.Features
	}
	orderPrint += "\n"
	orderPrint += "Количество: " + fmt.Sprintln(order.Amount)
	orderPrint += "Количество цветов: "
	if order.Cols == 0{
		orderPrint += "Я не знаю\n"
	} else if order.Cols == -1 {
		orderPrint += "Растровое изображение\n"
	} else {
		fmt.Sprintln(order.Cols)
	}
	orderPrint += "Срок: " + fmt.Sprintln(order.Deadline)
	orderPrint += "Комментарий: " + fmt.Sprintln(order.Comment)
	return orderPrint
}
