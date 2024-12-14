package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "goonline"
)

type Customer struct {
	customer_id int
	name        string
	email       string
	city        string
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dbConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbConnection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	// Select only one row
	var customer Customer
	err = db.QueryRow(`SELECT customer_id, name, email, city
					FROM customers 
					WHERE customer_id = $1`, 2).
		Scan(&customer.customer_id,
			&customer.name, &customer.email, &customer.city)
	handleError(err)
	fmt.Println(customer)

	// Insert only one row
	_, err = db.Exec(`INSERT INTO customers (name, email, city)
		VALUES ($1, $2, $3)`, "John", "john@mail.ru", "Hamburg")
	handleError(err)

	// Update row with id = 11
	_, err = db.Exec(`UPDATE customers 
				SET name = $1 WHERE customer_id = $2`, "Thomas", 11)
	handleError(err)

	// Delete row with id = 11
	_, err = db.Exec(`DELETE FROM customers 
						WHERE customer_id = $1`, 11)
	handleError(err)

	// Select all rows
	var customers []Customer
	rows, e := db.Query(`SELECT customer_id, name, email, city
					FROM customers`)
	handleError(e)
	defer rows.Close()
	for rows.Next() {
		var customer Customer
		rows.Scan(&customer.customer_id, &customer.name, &customer.email, &customer.city)
		customers = append(customers, customer)
	}

	fmt.Println(customers)
}
