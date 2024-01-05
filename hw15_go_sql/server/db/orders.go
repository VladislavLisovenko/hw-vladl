package db

import (
	"time"
)

func AddOrder(userID int, orderDate time.Time, totalAmount float64) (int, error) {
	lastInsertedID := 0
	queryString := "INSERT INTO public.\"Orders\"(user_id, order_date, total_amount) VALUES ($1, $2, $3) returning id"
	row := database.QueryRow(queryString, userID, orderDate, totalAmount)
	err := row.Scan(&lastInsertedID)
	if err != nil {
		return 0, err
	}

	return lastInsertedID, nil
}

func UpdateOrder(id int, userID int, orderDate time.Time, totalAmount float64) error {
	queryString := "update public.\"Orders\" set user_id=$1, order_date=$2, total_amount=$3 where id=$4"
	_, err := database.Exec(queryString, userID, orderDate, totalAmount, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOrder(id int) error {
	queryString := "delete from public.\"Orders\" where id=$1"
	_, err := database.Exec(queryString, id)
	if err != nil {
		return err
	}

	return nil
}
