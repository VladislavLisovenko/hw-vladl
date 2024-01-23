package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/db"
	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/entities"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	user := decodeEntity[entities.User](w, r)

	userID, err := db.AddUser(user)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Add("id", strconv.Itoa(userID))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := decodeEntity[entities.User](w, r)

	err := db.UpdateUser(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := decodeEntityID(w, r)

	err := db.DeleteUser(userID)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func UserList(w http.ResponseWriter, _ *http.Request) {
	userList, err := db.UserList()
	if err != nil {
		fmt.Println(err.Error())
	}
	userListDecoded, err := json.Marshal(userList)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Write(userListDecoded)
}
