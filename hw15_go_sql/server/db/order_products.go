package db

import (
	"context"
)

func AddOrderProduct(orderID int, productID int) (int, error) {
	lastInsertedID := 0
	queryString := "INSERT INTO public.\"OrderProducts\"(order_id, product_id) VALUES ($1, $2) returning id"
	row := database.QueryRow(queryString, orderID, productID)
	err := row.Scan(&lastInsertedID)
	if err != nil {
		return 0, err
	}

	return lastInsertedID, nil
}

func UpdateOrderProduct(id int, orderID int, productID int) error {
	tx, err := database.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	queryString := "update public.\"OrderProducts\" set order_id=$1, product_id=$2 where id=$3"
	_, err = database.Exec(queryString, orderID, productID, id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func DeleteOrderProduct(id int) error {
	queryString := "delete from public.\"OrderProducts\" where id=$1"
	_, err := database.Exec(queryString, id)
	if err != nil {
		return err
	}

	return nil
}
