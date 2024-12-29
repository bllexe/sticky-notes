package model

import "time"

type Color string

const (
	Yellow Color = "yellow"
	Blue   Color = "blue"
	Green  Color = "green"
	Pink   Color = "pink"
	Orange Color = "orange"
)

type Note struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Color     Color     `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
