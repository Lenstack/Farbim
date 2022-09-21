package entities

import "time"

const CategoryTableName = "categories"

type Category struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
