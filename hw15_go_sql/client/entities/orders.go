package entities

import "time"

type Order struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userId"`
	OrderDate   time.Time `json:"orderDate"`
	TotalAmount float64   `json:"totalAmount"`
}

func (o *Order) GetID() int {
	return o.ID
}

func (o *Order) SetID(id int) {
	o.ID = id
}
