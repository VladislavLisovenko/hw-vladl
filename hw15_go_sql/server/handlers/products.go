package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/db"
	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/entities"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	product := decodeEntity[entities.Product](w, r)

	productID, err := db.AddProduct(product)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Add("id", strconv.Itoa(productID))
}

func ProductList(w http.ResponseWriter, _ *http.Request) {
	productList, err := db.ProductList()
	if err != nil {
		fmt.Println(err.Error())
	}
	productListDecoded, err := json.Marshal(productList)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Write(productListDecoded)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := decodeEntityID(w, r)

	err := db.DeleteProduct(productID)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	product := decodeEntity[entities.Product](w, r)

	err := db.UpdateProduct(product)
	if err != nil {
		fmt.Println(err.Error())
	}
}
