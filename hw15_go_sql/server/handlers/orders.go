package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/db"
	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/entities"
)

func OrderList(w http.ResponseWriter, _ *http.Request) {
	orderList, err := db.OrderList()
	if err != nil {
		fmt.Println(err.Error())
	}
	orderListDecoded, err := json.Marshal(orderList)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Write(orderListDecoded)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	order := decodeEntity[entities.Order](w, r)

	err := db.UpdateOrder(order)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	order := decodeEntity[entities.Order](w, r)

	orderID, err := db.AddOrder(order)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Add("id", strconv.Itoa(orderID))
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := decodeEntityID(w, r)

	err := db.DeleteOrder(orderID)
	if err != nil {
		fmt.Println(err.Error())
	}
}
