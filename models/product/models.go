package models

import support "store/support"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}

func (c *Product) SetCreatedAt() {
	c.CreatedAt = support.CurrentDateTime()
}

func (c *Product) SetUpdatedAt() {
	c.UpdatedAt = support.CurrentDateTime()
}

func (c *Product) SetDeletedAt() {
	c.DeletedAt = support.CurrentDateTime()
}
