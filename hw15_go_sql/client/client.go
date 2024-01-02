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

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/client/entities"
)

const (
	ActionAdd    = "add"
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionSelect = "select"
)

type Entity interface {
	GetID() int
	SetID(int)
	Type() string
}

func sendRequest(url string, method string, action string, entity string, message []byte) (*http.Response, error) {
	getRequest, err := http.NewRequestWithContext(context.Background(), method, url, bytes.NewReader(message))
	if err != nil {
		return nil, err
	}
	getRequest.Header.Add("action", action)
	getRequest.Header.Add("entity", entity)

	httpClient := http.Client{}
	response, err := httpClient.Do(getRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func processEntity(url string, action string, entity Entity) {
	encMessage, err := json.Marshal(entity)
	if err != nil {
		log.Fatalln(err.Error())
	}

	response, err := sendRequest(url, http.MethodPost, action, entity.Type(), encMessage)
	if err != nil {
		log.Fatalf("%T %s error: %s", entity, action, err.Error())
	}
	defer response.Body.Close()

	if action == ActionAdd {
		id, err1 := strconv.Atoi(response.Header.Get("id"))
		if err1 != nil {
			fmt.Println(err1.Error())
		}
		entity.SetID(id)
	}
}

func main() {
	url := ""
	port := ""
	flag.StringVar(&url, "url", "http://localhost", "URL to send to")
	flag.StringVar(&port, "port", "8080", "Port to send to")
	flag.Parse()

	// user
	user := &entities.User{
		Name:     "Bob",
		Email:    "bob@mail.ru",
		Password: "123",
	}
	url += ":" + port
	processEntity(url, ActionAdd, user)

	user.Email = strings.ReplaceAll(user.Email, "mail.ru", "gmail.com")
	processEntity(url, ActionUpdate, user)

	// // products
	// products := []*entities.Product{
	// 	{
	// 		Name:  "Computer",
	// 		Price: 123456.78,
	// 	},
	// 	{
	// 		Name:  "Car",
	// 		Price: 12345678.90,
	// 	},
	// 	{
	// 		Name:  "Glasses",
	// 		Price: 1234.56,
	// 	},
	// }
	// for _, p := range products {
	// 	processEntity(url, ActionAdd, p)

	// 	p.Price /= 10
	// 	processEntity(url, ActionUpdate, p)
	// }

	// // order
	// order := &entities.Order{
	// 	UserID:      user.GetID(),
	// 	OrderDate:   time.Now(),
	// 	TotalAmount: 500000,
	// }
	// processEntity(url, ActionAdd, order)

	// order.OrderDate = time.Date(2024, time.April, 22, 5, 0, 0, 0, time.Local)
	// processEntity(url, ActionUpdate, order)

	// // orderProducts
	// orderProducts := []*entities.OrderProducts{}
	// for _, p := range products {
	// 	orderProducts = append(orderProducts, &entities.OrderProducts{
	// 		OrderID:   order.GetID(),
	// 		ProductID: p.GetID(),
	// 	})
	// }
	// for _, orderProduct := range orderProducts {
	// 	processEntity(url, ActionAdd, orderProduct)
	// }

	// // data removing
	// for _, orderProduct := range orderProducts {
	// 	processEntity(url, ActionDelete, orderProduct)
	// }

	// processEntity(url, ActionDelete, order)

	// for _, p := range products {
	// 	processEntity(url, ActionDelete, p)
	// }

	// processEntity(url, ActionDelete, user)
}
