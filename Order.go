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

func initOrder() *Order {
	return &Order{
		Orderid:     uuid.New(),
		Type:        "",
		ProductName: "",
		Features:    "",
		Amount:      0,
		Cols:        0,
		Mockup:      "",
		Layout:      "",
		CustomerID:  "",
		Deadline:    "",
		Comment:     "",
		State:       "new",
		edit:        false,
		state:       0,
	}
}

func stringifyOrder(order *Order) string {
	var orderPrint string
	orderPrint += "Тип заказа: " + fmt.Sprintln(l10n[order.Type])
	orderPrint += "Изделие: " + fmt.Sprintln(l10n[order.ProductName])
	if len(order.Features) != 0 {
		orderPrint += "Особенности: " + order.Features
	}
	orderPrint += "Количество: " + fmt.Sprintln(order.Amount)
	orderPrint += "Количество цветов: " + fmt.Sprintln(order.Cols)
	orderPrint += "Срок: " + fmt.Sprintln(order.Deadline)
	orderPrint += "Комментарий: " + fmt.Sprintln(order.Comment)
	return orderPrint
}
