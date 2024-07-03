package models

type Item struct {
	Id       string
	Name     string
	Quantity int
	Type     string
}

var Bag = []Item{}
