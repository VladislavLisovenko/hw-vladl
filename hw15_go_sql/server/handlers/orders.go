package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/db"
	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/entities"
)

func AddOrder(w http.ResponseWriter, r *http.Request) {
	entity := decodeEntity[entities.Order](w, r)

	id, err := db.AddOrder(entity)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Add("id", strconv.Itoa(id))
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	entity := decodeEntity[entities.Order](w, r)

	err := db.UpdateOrder(entity)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	entityID := decodeEntityID[entities.Order](w, r)

	err := db.DeleteOrder(entityID)
	if err != nil {
		fmt.Println(err.Error())
	}
}
