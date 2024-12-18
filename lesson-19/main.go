package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
}

type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
}

type Order struct {
	OrderID     int
	CustomerId  int
	ProductID   int
	OrderDate   time.Time
	TotalAmount int
}

type Customer struct {
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "goonline"
)

func main() {
	dbConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbConnection)
	handleError(err)

	tx, err := db.Begin()
	handleError(err)
	defer tx.Commit()

	newProduct := Product{11, "Iphone", "Electronics", 1200}
	_, err = tx.Exec(`INSERT INTO products VALUES ($1, $2, $3, $4)`,
		newProduct.ID, newProduct.Name,
		newProduct.Category, newProduct.Price)

	handleError(err)
	newOrder := Order{
		CustomerId:  4,
		ProductID:   newProduct.ID,
		OrderDate:   time.Date(2024, 12, 17, 0, 0, 0, 0, time.UTC),
		TotalAmount: 10,
	}
	_, err = tx.Exec(`INSERT INTO orders VALUES ($1, $2, $3, $4)`,
		newOrder.CustomerId, newOrder.ProductID,
		newOrder.OrderDate, newOrder.TotalAmount,
	)
	handleError(err)
}
