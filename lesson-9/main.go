// package main

// import "fmt"

// type Messenger interface {
// 	send(msg string, receiverId int)
// }
// type Telegram struct {
// 	name string
// }
// type Whatsapp struct {
// 	name string
// }

// func (t Telegram) send(msg string, receiverId int) {
// 	fmt.Println("Sent via Telegram")
// }

// func (t Whatsapp) send(msg string, receiverId int) {
// 	fmt.Println("Sent via Whatsapp")
// }

// func share(m Messenger, msg string, receiverID int) {
// 	m.send(msg, receiverID)
// }

// func main() {
// 	telegram := Telegram{"Telegram"}
// 	whatsapp := Whatsapp{"Whatsapp"}

// 	share(telegram, "Hello from Telegram", 1)
// 	share(whatsapp, "Hello from Whatsapp", 1)
// }
