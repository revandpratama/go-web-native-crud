package entities

import "time"

type Category struct {
	Id uint
	Name string
	Slug string
	CreatedAt time.Time
	UpdatedAt time.Time
}
 