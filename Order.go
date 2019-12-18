package main

import "github.com/google/uuid"

type Order struct {
	id          uuid.UUID `db: orderid`
	Type        string    `db: type`
	ProductName string    `db: productname`
	Features    []string  `db: Features`
	Amount      int       `db: Amount`
	Cols        int       `db: cols`
	Mockup      string    `db: Mockup`
	Layout      string    `db: Layout`
	CustomerID  string    `db: customerid`
	Deadline    string    `db: Deadline`
	Comment     string    `db: Comment`
	State       string    `db: State`
	edit        bool
	state 		int
}

var orders = make(map[int64]*Order)

func initOrder() *Order {
	return &Order{
		id:          uuid.New(),
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
		state: 		0,
	}
}