package main

import (
	"context"
	"github.com/jackc/pgx"
	"log"
)

func initConnection() *pgx.Conn{
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:@localhost:5432/nani")
	if err != nil{
		log.Fatal(err)
	}
	_, err = conn.Exec(context.Background(),"CREATE SCHEMA IF NOT EXISTS orders")

	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS nani.orders.customers(" +
		"CustomerID int primary key," +
		"TelegramUsername varchar not null)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(context.Background(),"CREATE TABLE IF NOT EXISTS nani.orders.orders(" +
		"OrderID int primary key," +
		"Type varchar NOT NULL," +
		"ProductName varchar NOT NULL," +
		"Features varchar not null," +
		"Amount int not null," +
		"Cols int not null," +
		"Layout varchar not null," +
		"Mockup varchar not null, " +
		"Deadline varchar not null," +
		"CustomerID int  FOREIGN KEY REFERENCES  nani.orders.customers(CustomerID))" )

	if err != nil {
		log.Fatal(err)
	}
		return conn
}



func getNewOrders(){

}

func getInProgresOrders(){

}

func getDoneOrders(){

}

func moveNewToInProgress(id int64){

}

func moveInProgressToDone(id int64){

}

func putOrder(order Order, id int64){

}
