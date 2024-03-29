package db

import "github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/entities"

func AddProduct(entity entities.Product) (int, error) {
	name := entity.Name
	price := entity.Price

	lastInsertedID := 0
	queryString := `
	INSERT INTO public."Products"(name, price)
	VALUES ($1, $2) 
	returning id`
	row := database.QueryRow(queryString, name, price)
	err := row.Scan(&lastInsertedID)
	if err != nil {
		return 0, err
	}

	return lastInsertedID, nil
}

func UpdateProduct(entity entities.Product) error {
	id := entity.GetID()
	name := entity.Name
	price := entity.Price

	queryString := `
	update public."Products"
	set name=$1, price=$2 
	where id=$3`
	_, err := database.Exec(queryString, name, price, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(id int) error {
	queryString := `
	delete from public."Products"
	where id=$1`
	_, err := database.Exec(queryString, id)
	if err != nil {
		return err
	}

	return nil
}

func ProductList() ([]entities.Product, error) {
	queryString := `
	SELECT id, name, price 
	FROM public."Products"`
	rows, err := database.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productList := make([]entities.Product, 0)
	for rows.Next() {
		var id int
		var name string
		var price float64
		if err = rows.Scan(&id, &name, &price); err != nil {
			return nil, err
		}
		productList = append(productList, entities.Product{
			ID:    id,
			Name:  name,
			Price: price,
		},
		)
	}

	return productList, nil
}
