package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/db"
	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/entities"
	"github.com/go-chi/chi"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var entity entities.Product
	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	id, err := db.AddProduct(entity.Name, entity.Price)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Add("id", strconv.Itoa(id))
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var entity entities.Product
	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	err = db.UpdateProduct(entity.GetID(), entity.Name, entity.Price)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	entityID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	err = db.DeleteProduct(entityID)
	if err != nil {
		fmt.Println(err.Error())
	}
}
