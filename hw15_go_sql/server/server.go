package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/handlers"
	"github.com/go-chi/chi"
)

func main() {
	url := ""
	port := ""
	flag.StringVar(&url, "url", "localhost", "URL to listen from, without protocol, e.g. 'localhost'")
	flag.StringVar(&port, "port", "8080", "Port to listen from")
	flag.Parse()

	router := chi.NewRouter()

	router.Post("/users", handlers.AddUser)
	router.Post("/users/{id:[0-9]+}", handlers.UpdateUser)
	router.Delete("/users/{id:[0-9]+}", handlers.DeleteUser)

	router.Post("/products", handlers.AddProduct)
	router.Post("/products/{id:[0-9]+}", handlers.UpdateProduct)
	router.Delete("/products/{id:[0-9]+}", handlers.DeleteProduct)

	router.Post("/orders", handlers.AddOrder)
	router.Post("/orders/{id:[0-9]+}", handlers.UpdateOrder)
	router.Delete("/orders/{id:[0-9]+}", handlers.DeleteOrder)

	router.Post("/order_products", handlers.AddOrderProduct)
	router.Post("/order_products/{id:[0-9]+}", handlers.UpdateOrderProduct)
	router.Delete("/order_products/{id:[0-9]+}", handlers.DeleteOrderProduct)

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
