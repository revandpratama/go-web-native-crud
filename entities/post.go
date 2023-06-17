package entities

import "time"

type Post struct {
	Id          uint
	User_id     uint
	Category    Category
	Title       string
	Slug        string
	Image string
	Excerpt     string
	Body        string
	PublishedAt string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
