package models

type Bag struct {
	Id       string
	Name     string
	Quantity int
	Type     string
}

var Inventory = []Bag{}
