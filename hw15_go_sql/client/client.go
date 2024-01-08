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
	flag.StringVar(&port, "port", "8080", "Port to send to")
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

	// order
	order := &entities.Order{
		UserID:    user.GetID(),
		OrderDate: time.Now(),
		Products:  products,
	}
	addr = fmt.Sprintf("%s/%s", url, "orders")
	processEntity(addr, http.MethodPost, order)

	for _, p := range products {
		p.Price /= 10
	}
	order.OrderDate = time.Date(2024, time.April, 22, 5, 0, 0, 0, time.Local)
	addr = fmt.Sprintf("%s/%s/%d", url, "orders", order.GetID())
	processEntity(addr, http.MethodPost, order)

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
