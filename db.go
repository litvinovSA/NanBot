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
		id SERIAL primary key,
		OrderID UUID not null,
		Type text NOT NULL,
		ProductName text NOT NULL,
		Features text,
		Amount int not null,
		Cols int not null,
		Layout text not null,
		Mockup text not null,
		Deadline text not null,
		State text not null,
		Comment text not null,
		CustomerID int REFERENCES  orders.customers(CustomerID))
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
	var newOrders []Order
	err := db.Select(&newOrders, ".. WHERE State = $1", "new")
	if err != nil{
		log.Fatal(err)
	}
	return newOrders
}

func getInProgresOrders(db *sqlx.DB) []Order{
	var newOrders []Order
	err := db.Select(&newOrders, ".. WHERE State = $1", "inprogress")
	if err != nil{
		log.Fatal(err)
	}
	return newOrders
}

func getDoneOrders(db *sqlx.DB) []Order{
	var newOrders []Order
	err := db.Select(&newOrders, ".. WHERE State = $1", "done")
	if err != nil{
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
	fields := "type, productname,  Amount, cols, Layout, Mockup, Deadline, State, Comment, orderid"
	fmt.Println(&order.CustomerID)

	_, err := db.Exec("INSERT into customers (TelegramUsername) values ($1)", order.CustomerID)
	if err != nil {
		log.Println("Existing customer: ", err)
	}
	res, err := db.Exec("INSERT INTO orders ("+ fields +") values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		order.Type,
		order.ProductName,
		order.Amount,
		order.Cols,
		order.Layout,
		order.Mockup,
		order.Deadline,
		order.State,
		order.Comment,
		order.id)
	if err != nil{
		log.Fatal(err)
	}
	res.LastInsertId()
}
