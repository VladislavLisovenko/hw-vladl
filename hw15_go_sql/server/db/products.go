package db

func AddProduct(name string, price float64) (int, error) {
	lastInsertedID := 0
	row := database.QueryRow("INSERT INTO public.\"Products\"(name, price) VALUES ($1, $2) returning id", name, price)
	err := row.Scan(&lastInsertedID)
	if err != nil {
		return 0, err
	}

	return lastInsertedID, nil
}

func UpdateProduct(id int, name string, price float64) error {
	_, err := database.Exec("update public.\"Products\" set name=$1, price=$2 where id=$3", name, price, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(id int) error {
	_, err := database.Exec("delete from public.\"Products\" where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
