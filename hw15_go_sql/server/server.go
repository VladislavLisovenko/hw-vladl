package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/client/entities"
	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/db"
)

const (
	ActionAdd    = "add"
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionSelect = "select"
)

const (
	EntityUser         = "user"
	EntityProduct      = "product"
	EntityOrder        = "order"
	EntityOrderProduct = "order_product"
)

type Entity interface {
	GetID() int
	SetID(int)
	Type() string
}

// func decodedEntity[T Entity](body io.ReadCloser, entity T) error {
// 	err := json.NewDecoder(body).Decode(entity)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func decodedUser(body io.ReadCloser) (entities.User, error) {
	var entity entities.User
	err := json.NewDecoder(body).Decode(&entity)
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func decodedProduct(body io.ReadCloser) (entities.Product, error) {
	var entity entities.Product
	err := json.NewDecoder(body).Decode(&entity)
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func decodedOrder(body io.ReadCloser) (entities.Order, error) {
	var entity entities.Order
	err := json.NewDecoder(body).Decode(&entity)
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func decodedOrderProduct(body io.ReadCloser) (entities.OrderProducts, error) {
	var entity entities.OrderProducts
	err := json.NewDecoder(body).Decode(&entity)
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("action")

	switch action {
	case ActionAdd:
		entity, err := decodedUser(r.Body)
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
	case ActionUpdate:
		entity, err := decodedUser(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		err = db.UpdateUser(entity.ID, entity.Name, entity.Email, entity.Password)
		if err != nil {
			fmt.Println(err.Error())
		}
	case ActionDelete:
		entity, err := decodedUser(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		err = db.DeleteUser(entity.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleProduct(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("action")
	switch action {
	case ActionAdd:
		entity, err := decodedProduct(r.Body)
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
	case ActionUpdate:
		entity, err := decodedProduct(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		err = db.UpdateProduct(entity.ID, entity.Name, entity.Price)
		if err != nil {
			fmt.Println(err.Error())
		}
	case ActionDelete:
		entity, err := decodedProduct(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		err = db.DeleteProduct(entity.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("action")
	switch action {
	case ActionAdd:
		entity, err := decodedOrder(r.Body)
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
	case ActionUpdate:
		entity, err := decodedOrder(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		err = db.UpdateOrder(entity.ID, entity.UserID, entity.OrderDate, entity.TotalAmount)
		if err != nil {
			fmt.Println(err.Error())
		}
	case ActionDelete:
		entity, err := decodedOrder(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		err = db.DeleteOrder(entity.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleOrderProduct(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("action")
	switch action {
	case ActionAdd:
		entity, err := decodedOrderProduct(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		id, err := db.AddOrderProduct(entity.OrderID, entity.ProductID)
		if err != nil {
			fmt.Println(err.Error())
		}
		w.Header().Add("id", strconv.Itoa(id))
	case ActionUpdate:
		entity, err := decodedOrderProduct(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		err = db.UpdateOrderProduct(entity.ID, entity.OrderID, entity.ProductID)
		if err != nil {
			fmt.Println(err.Error())
		}
	case ActionDelete:
		entity, err := decodedOrderProduct(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		err = db.DeleteOrderProduct(entity.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	url := ""
	port := ""
	flag.StringVar(&url, "url", "localhost", "URL to listen from, without protocol, e.g. 'localhost'")
	flag.StringVar(&port, "port", "8080", "Port to listen from")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		entity := r.Header.Get("entity")
		switch entity {
		case EntityUser:
			handleUser(w, r)
		case EntityProduct:
			handleProduct(w, r)
		case EntityOrder:
			handleOrder(w, r)
		case EntityOrderProduct:
			handleOrderProduct(w, r)
		}
	})

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	srv.Addr = fmt.Sprintf("%s:%s", url, port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
