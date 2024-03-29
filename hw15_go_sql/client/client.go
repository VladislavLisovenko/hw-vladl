package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/client/entities"
)

func sendRequest(url string, method string, message []byte) (*http.Response, error) {
	getRequest, err := http.NewRequestWithContext(context.Background(), method, url, bytes.NewReader(message))
	if err != nil {
		return nil, err
	}

	httpClient := http.Client{}
	response, err := httpClient.Do(getRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func processEntity[T entities.HasID](url string, httpMethod string, entity T) {
	encMessage, err := json.Marshal(entity)
	if err != nil {
		log.Fatalln(err.Error())
	}

	response, err := sendRequest(url, httpMethod, encMessage)
	if err != nil {
		log.Fatalf("%T %s error: %s", entity, httpMethod, err.Error())
	}
	defer response.Body.Close()

	headerID := response.Header.Get("id")
	if headerID != "" {
		id, err1 := strconv.Atoi(headerID)
		if err1 != nil {
			fmt.Println(err1.Error())
		}
		entity.SetID(id)
	}
}

func processEntityList[T []entities.User | []entities.Product | []entities.Order](url string) T {
	response, err := sendRequest(url, http.MethodGet, nil)
	if err != nil {
		log.Fatalf("%s on GET error: %s", url, err.Error())
	}
	defer response.Body.Close()

	var entity T
	err = json.NewDecoder(response.Body).Decode(&entity)
	if err != nil {
		fmt.Printf("%s on GET error: %s", url, err.Error())
	}

	return entity
}

func processEntityRemoving(url string) {
	response, err := sendRequest(url, http.MethodDelete, nil)
	if err != nil {
		log.Fatalf("%s on DELETE error: %s", url, err.Error())
	}
	defer response.Body.Close()
}

func main() {
	url := ""
	port := ""
	flag.StringVar(&url, "url", "http://localhost", "URL to send to")
	flag.StringVar(&port, "port", "5436", "Port to send to")
	flag.Parse()

	url += ":" + port

	// user
	user := &entities.User{
		Name:     "Bob",
		Email:    "bob@mail.ru",
		Password: "123",
	}

	addr := fmt.Sprintf("%s/%s", url, "users")
	processEntity(addr, http.MethodPost, user)

	user.Email = strings.ReplaceAll(user.Email, "mail.ru", "gmail.com")
	addr = fmt.Sprintf("%s/%s/%d", url, "users", user.GetID())
	processEntity(addr, http.MethodPost, user)

	addr = fmt.Sprintf("%s/%s", url, "users")
	userList := processEntityList[[]entities.User](addr)
	for _, u := range userList {
		fmt.Println(u)
	}

	// products
	addr = fmt.Sprintf("%s/%s", url, "products")
	products := []*entities.Product{
		{
			Name:  "Computer",
			Price: 123456.78,
		},
		{
			Name:  "Car",
			Price: 12345678.90,
		},
		{
			Name:  "Glasses",
			Price: 1234.56,
		},
	}
	for _, p := range products {
		processEntity(addr, http.MethodPost, p)
	}

	addr = fmt.Sprintf("%s/%s", url, "products")
	productList := processEntityList[[]entities.Product](addr)
	for _, p := range productList {
		fmt.Println(p)
	}

	// order
	order := &entities.Order{
		UserID:    user.GetID(),
		OrderDate: time.Now(),
		Products:  products,
	}
	addr = fmt.Sprintf("%s/%s", url, "orders")
	processEntity(addr, http.MethodPost, order)

	order.OrderDate = time.Date(2024, time.April, 22, 5, 0, 0, 0, time.Local)
	addr = fmt.Sprintf("%s/%s/%d", url, "orders", order.GetID())
	processEntity(addr, http.MethodPost, order)

	addr = fmt.Sprintf("%s/%s", url, "orders")
	orderList := processEntityList[[]entities.Order](addr)
	for _, o := range orderList {
		fmt.Println(o)
	}

	// data removing
	addr = fmt.Sprintf("%s/%s/%d", url, "orders", order.GetID())
	processEntityRemoving(addr)

	for _, p := range products {
		addr = fmt.Sprintf("%s/%s/%d", url, "products", p.GetID())
		processEntityRemoving(addr)
	}

	addr = fmt.Sprintf("%s/%s/%d", url, "users", user.GetID())
	processEntityRemoving(addr)
}
