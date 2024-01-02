package db

func AddUser(name string, email string, password string) (int, error) {
	lastInsertedID := 0
	queryString := "INSERT INTO public.\"Users\"(name, email, password) VALUES ($1, $2, $3) returning id"
	row := database.QueryRow(queryString, name, email, password)
	err := row.Scan(&lastInsertedID)
	if err != nil {
		return 0, err
	}

	return lastInsertedID, nil
}

func UpdateUser(id int, name string, email string, password string) error {
	queryString := "update public.\"Users\" set name=$1, email=$2, password=$3 where id=$4"
	_, err := database.Exec(queryString, name, email, password, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	queryString := "delete from public.\"Users\" where id=$1"
	_, err := database.Exec(queryString, id)
	if err != nil {
		return err
	}

	return nil
}
