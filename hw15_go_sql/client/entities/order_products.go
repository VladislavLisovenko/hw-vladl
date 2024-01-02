package entities

type OrderProducts struct {
	ID        int `json:"id"`
	OrderID   int `json:"orderId"`
	ProductID int `json:"productId"`
}

func (o *OrderProducts) GetID() int {
	return o.ID
}

func (o *OrderProducts) SetID(id int) {
	o.ID = id
}

func (o *OrderProducts) Type() string {
	return "order_product"
}
