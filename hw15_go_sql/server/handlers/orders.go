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

func AddOrder(w http.ResponseWriter, r *http.Request) {
	var entity entities.Order
	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	id, err := db.AddOrder(entity.UserID, entity.OrderDate, entity.TotalAmount)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Add("id", strconv.Itoa(id))
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var entity entities.Order
	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	err = db.UpdateOrder(entity.GetID(), entity.UserID, entity.OrderDate, entity.TotalAmount)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	entityID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	err = db.DeleteOrder(entityID)
	if err != nil {
		fmt.Println(err.Error())
	}
}
