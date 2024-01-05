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

func AddUser(w http.ResponseWriter, r *http.Request) {
	var entity entities.User
	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	id, err := db.AddUser(entity.Name, entity.Email, entity.Password)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Add("id", strconv.Itoa(id))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var entity entities.User
	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	err = db.UpdateUser(entity.GetID(), entity.Name, entity.Email, entity.Password)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	entityID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	err = db.DeleteUser(entityID)
	if err != nil {
		fmt.Println(err.Error())
	}
}
