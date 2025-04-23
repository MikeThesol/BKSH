package models

type Book struct {
	ID          uint64
	Author      string
	Name        string
	Quotes      []string // цитаты
	Score       uint8    // оценка
	Description string
	PhotoURL    string
	Private     bool
}
