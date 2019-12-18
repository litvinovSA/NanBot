package main

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

var schema = `
		CREATE SCHEMA IF NOT EXISTS orders;

		CREATE TABLE IF NOT EXISTS customers(
		CustomerID SERIAL primary key,
		TelegramUsername text not null unique);
		
		CREATE TABLE IF NOT EXISTS orders(
		"id" SERIAL primary key,
		"orderid" UUID not null,
		"type" text NOT NULL,
		"productname" text NOT NULL,
		"amount" int not null,
		"cols" int not null,
		"layout" text not null,
		"mockup" text not null,
		"deadline" text not null,
		"state" text not null,
		"comment" text not null,
		"customerid" int )
`

func initConnection() *sqlx.DB {
	db, err := sqlx.Connect("pgx", "postgres://postgres:@localhost:5432/nani")
	if err != nil {
		log.Fatal(err)
	}
	db.MustExec(schema)
	return db
}

func getNewOrders(db *sqlx.DB) []Order{
	newOrders := []Order{}
	err := db.Select(&newOrders, "SELECT * FROM orders WHERE state=$1", "new")
	if err != nil {
		log.Fatal(err)
	}
	return newOrders
}

func getInProgresOrders(db *sqlx.DB) []Order {
	var newOrders []Order
	err := db.Select(&newOrders, "SELECT * FROM orders WHERE state=$1", "inprogress")
	if err != nil {
		log.Fatal(err)
	}
	return newOrders
}

func getDoneOrders(db *sqlx.DB) []Order {
	var newOrders []Order
	err := db.Select(&newOrders, "SELECT * FROM orders WHERE state=$1", "done")
	if err != nil {
		log.Fatal(err)
	}
	return newOrders
}

func moveNewToInProgress(id int64) {

}

func moveInProgressToDone(id int64) {

}

func putOrder(order Order, db *sqlx.DB) {
	fmt.Println(printOrder(&order))
	fields := "type, productname,  Amount, cols, Layout, Mockup, Deadline, State, Comment, orderid, customerid"
	fmt.Println(&order.CustomerID)

	_, err := db.Exec("INSERT into customers (TelegramUsername) values ($1)", order.CustomerID)
	if err != nil {
		log.Println("Existing customer: ", err)
	}
	res, err := db.Exec("INSERT INTO orders ("+fields+") values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		order.Type,
		order.ProductName,
		order.Amount,
		order.Cols,
		order.Layout,
		order.Mockup,
		order.Deadline,
		order.State,
		order.Comment,
		order.Orderid,
		order.CustomerID,
)
	if err != nil {
		log.Fatal(err)
	}
	res.LastInsertId()
}
