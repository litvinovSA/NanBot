package main

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
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
		"customerid" text )
`

func initConnection() *sqlx.DB {
	db, err := sqlx.Connect("pgx", "postgres://postgres:@localhost:5432/nani")
	if err != nil {
		log.Fatal(err)
	}
	db.MustExec(schema)
	return db
}

func getNewOrders(db *sqlx.DB) []Order {
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

func moveToProgress(id int, db *sqlx.DB) {
	_, err := db.Exec("UPDATE orders SET state = $1 WHERE id = $2", "inprogress", strconv.Itoa(id))
	if err != nil {
		log.Fatal(err)
	}
}

func moveToDone(id int, db *sqlx.DB) {
	_, err := db.Exec("UPDATE orders SET state = $1 WHERE id = $2", "done", strconv.Itoa(id))
	if err != nil {
		log.Fatal(err)
	}
}
func getNextID(db *sqlx.DB) int {
	rows, err := db.Query("Select nextval(pg_get_serial_sequence('orders', 'id')) as new_id;")
	if err != nil {
		log.Fatal(err)
	}
	var id int
	for rows.Next() {
		rows.Scan(&id)
	}
	return id
}

func getUserOrders(db *sqlx.DB, username string) []Order {
	var newOrders []Order
	err := db.Select(&newOrders, "SELECT * FROM orders WHERE customerid=$1", username)
	if err != nil {
		log.Fatal(err)
	}
	return newOrders
}

func putOrder(order Order, db *sqlx.DB) {
	fields := "type, productname,  Amount, cols, Layout, Mockup, Deadline, State, Comment, orderid, customerid"

	_, err := db.Exec("INSERT into customers (TelegramUsername) values ($1)", order.CustomerID)
	if err != nil {
		log.Println("Existing customer: ", err)
	}
	_, err = db.Exec("INSERT INTO orders ("+fields+") values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
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
		log.Println(err)
	}
}
