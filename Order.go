package main

import (
	"github.com/google/uuid"
)

type Order struct {
	Id          int
	Orderid     uuid.UUID `db:"orderid"`
	Type        string    `db:"type"`
	ProductName string    `db:"productname"`
	Features    []string  `db:"features"`
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
		Features:    nil,
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
