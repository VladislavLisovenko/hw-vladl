package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/entities"
	"github.com/go-chi/chi"
)

type Entity interface {
	entities.User | entities.Product | entities.Order
}

func decodeEntity[T Entity](w http.ResponseWriter, r *http.Request) T {
	var entity T
	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return entity
}

func decodeEntityID(w http.ResponseWriter, r *http.Request) int {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return id
}
