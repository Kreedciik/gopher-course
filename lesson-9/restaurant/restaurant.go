package restaurant

import (
	"fmt"
	"time"
)

type OrderManager interface {
	PlaceOrder(string) string
	CancelOrder(int) string
	OrderStatus(int) string
}

const (
	PENDING   = "Pending"
	DONE      = "Done"
	CANCELLED = "Cancelled"
)

// Order status -> Pending | Done | Cancelled

type DineIn struct {
	Type        string
	orderID     int
	orderStatus string
	TableNumber int
}

type TakeAway struct {
	Type        string
	orderID     int
	orderStatus string
}

type Delivery struct {
	Type        string
	orderID     int
	orderStatus string
	Address     string
}

func (d *DineIn) PlaceOrder(order string) string {
	id := time.Now().Second()
	fmt.Printf("The order with #%d created for table number %d \n", id, d.TableNumber)
	d.orderStatus = PENDING
	return order
}
func (d *DineIn) CancelOrder(orderID int) string {
	fmt.Printf("The order with #%d cancelled for table number %d \n", orderID, d.TableNumber)
	d.orderStatus = CANCELLED
	return "Cancelled"
}
func (d *DineIn) OrderStatus(orderID int) string {
	fmt.Printf("The status of order with #%d for table number %d is %s \n", orderID, d.TableNumber, d.orderStatus)
	d.orderStatus = CANCELLED
	return d.orderStatus
}

func (t *TakeAway) PlaceOrder(order string) string {
	id := time.Now().Second()
	fmt.Printf("The order with #%d created \n", id)
	t.orderStatus = PENDING
	return order
}
func (t *TakeAway) CancelOrder(orderID int) string {
	fmt.Printf("The order with #%d cancelled \n", orderID)
	t.orderStatus = CANCELLED
	return "Cancelled"
}
func (t *TakeAway) OrderStatus(orderID int) string {
	fmt.Printf("The status of order with ID #%d is %s \n", orderID, t.orderStatus)
	t.orderStatus = CANCELLED
	return t.orderStatus
}

func (d *Delivery) PlaceOrder(order string) string {
	id := time.Now().Second()
	fmt.Printf("The order with #%d created for delivering with address %s \n", id, d.Address)
	d.orderStatus = PENDING
	return order
}
func (d *Delivery) CancelOrder(orderID int) string {
	fmt.Printf("The order with #%d cancelled for delivering \n", orderID)
	d.orderStatus = CANCELLED
	return "Cancelled"
}
func (d *Delivery) OrderStatus(orderID int) string {
	fmt.Printf("The status of order with ID #%d is %s \n", orderID, d.orderStatus)
	d.orderStatus = CANCELLED
	return d.orderStatus
}
