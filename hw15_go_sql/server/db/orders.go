package db

import (
	"context"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/entities"
)

func orderTotalAmount(orderID int) (float64, error) {
	v := 0.0
	queryString := `
	SELECT SUM(P.PRICE) AS OrderTotalAmount
	FROM PUBLIC."Products" AS P
	INNER JOIN PUBLIC."OrderProducts" AS OP ON P.ID = OP.PRODUCT_ID
	AND OP.ORDER_ID = $1`
	row := database.QueryRow(queryString, orderID)
	err := row.Scan(&v)
	if err != nil {
		return 0.0, err
	}

	return v, nil
}

func AddOrder(entity entities.Order) (int, error) {
	userID := entity.UserID
	orderDate := entity.OrderDate
	products := entity.Products

	tx, err := database.BeginTx(context.Background(), nil)
	if err != nil {
		return 0, err
	}

	orderID := 0
	queryString := `
	INSERT INTO public."Orders"(user_id, order_date) 
	VALUES ($1, $2) returning id`
	row := database.QueryRow(queryString, userID, orderDate)
	err = row.Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	queryString = `
	INSERT INTO public."OrderProducts"(order_id, product_id) 
	VALUES ($1, $2)`
	for _, p := range products {
		_, err = database.Exec(queryString, orderID, p.GetID())
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	totalAmount, err := orderTotalAmount(orderID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	queryString = `
	UPDATE PUBLIC."Orders"
	SET TOTAL_AMOUNT = $1
	WHERE ID = $2`
	_, err = database.Exec(queryString, totalAmount, orderID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return orderID, nil
}

func UpdateOrder(entity entities.Order) error {
	id := entity.GetID()
	userID := entity.UserID
	orderDate := entity.OrderDate
	products := entity.Products

	tx, err := database.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	queryString := `DELETE FROM public."OrderProducts" WHERE order_id=$1`
	_, err = database.Exec(queryString, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryString = `
	INSERT INTO public."OrderProducts"(order_id, product_id) 
	VALUES ($1, $2)`
	for _, p := range products {
		_, err = database.Exec(queryString, id, p.GetID())
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	totalAmount, err := orderTotalAmount(id)
	if err != nil {
		tx.Rollback()
		return err
	}
	queryString = `
	UPDATE public."Orders" SET user_id=$1, order_date=$2, total_amount=$3
	WHERE id=$4`
	_, err = database.Exec(queryString, userID, orderDate, totalAmount, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func DeleteOrder(id int) error {
	queryString := `delete from public."Orders" where id=$1`
	_, err := database.Exec(queryString, id)
	if err != nil {
		return err
	}

	return nil
}
