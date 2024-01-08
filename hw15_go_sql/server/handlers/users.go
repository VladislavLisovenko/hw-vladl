package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/db"
	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/entities"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	entity := decodeEntity[entities.User](w, r)

	id, err := db.AddUser(entity)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Add("id", strconv.Itoa(id))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	entity := decodeEntity[entities.User](w, r)

	err := db.UpdateUser(entity)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	entityID := decodeEntityID[entities.User](w, r)

	err := db.DeleteUser(entityID)
	if err != nil {
		fmt.Println(err.Error())
	}
}
