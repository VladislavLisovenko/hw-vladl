package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/db"
	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/entities"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	entity := decodeEntity[entities.Product](w, r)

	id, err := db.AddProduct(entity)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Add("id", strconv.Itoa(id))
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	entity := decodeEntity[entities.Product](w, r)

	err := db.UpdateProduct(entity)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	entityID := decodeEntityID[entities.Product](w, r)

	err := db.DeleteProduct(entityID)
	if err != nil {
		fmt.Println(err.Error())
	}
}
