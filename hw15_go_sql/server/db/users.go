package db

import "github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/entities"

func AddUser(entity entities.User) (int, error) {
	name := entity.Name
	email := entity.Email
	password := entity.Password

	lastInsertedID := 0
	queryString := `
	INSERT INTO public."Users"(name, email, password)
	VALUES ($1, $2, $3) returning id`
	row := database.QueryRow(queryString, name, email, password)
	err := row.Scan(&lastInsertedID)
	if err != nil {
		return 0, err
	}

	return lastInsertedID, nil
}

func UpdateUser(entity entities.User) error {
	id := entity.GetID()
	name := entity.Name
	email := entity.Email
	password := entity.Password

	queryString := `
	update public."Users" 
	set name=$1, email=$2, password=$3 
	where id=$4`
	_, err := database.Exec(queryString, name, email, password, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	queryString := `delete from public."Users" where id=$1`
	_, err := database.Exec(queryString, id)
	if err != nil {
		return err
	}

	return nil
}

func UserList() ([]entities.User, error) {
	queryString := `
	SELECT id, name, email, password 
	FROM public."Users"`
	rows, err := database.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userList := make([]entities.User, 0)
	for rows.Next() {
		var userID int
		var name, email, password string
		if err = rows.Scan(&userID, &name, &email, &password); err != nil {
			return nil, err
		}
		userList = append(userList, entities.User{
			ID:       userID,
			Name:     name,
			Email:    email,
			Password: password,
		},
		)
	}

	return userList, nil
}
