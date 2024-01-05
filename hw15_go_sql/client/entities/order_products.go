package entities

type OrderProduct struct {
	ID        int `json:"id"`
	OrderID   int `json:"orderId"`
	ProductID int `json:"productId"`
}

func (o *OrderProduct) GetID() int {
	return o.ID
}

func (o *OrderProduct) SetID(id int) {
	o.ID = id
}
