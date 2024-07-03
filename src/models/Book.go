package models

type Book struct {
	Id     string
	Title  string
	Author string
}

var Books = []Book{
	{Id: "1", Title: "CSS50 Wet Dreams", Author: "FAANG"},
	{Id: "2", Title: "CSS50 Fear", Author: "Leetcode.co"},
	{Id: "3", Title: "CSS50 Reality", Author: "Linkedin"},
}
