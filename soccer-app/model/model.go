package model

type Player struct {
	Id       uint
	Name     string
	Position string
	Age      uint
}
type Team struct {
	Id      uint
	Name    string
	Country string
	Players []Player
}