package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/VladislavLisovenko/hw-vladl/hw15_go_sql/server/handlers"
	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Server is started...")

	url := ""
	port := ""
	flag.StringVar(&url, "url", "", "URL to listen from, without protocol, e.g. 'localhost'")
	flag.StringVar(&port, "port", "8080", "Port to listen from")
	flag.Parse()

	router := chi.NewRouter()

	router.Post("/users", handlers.AddUser)
	router.Get("/users", handlers.UserList)
	router.Post("/users/{id:[0-9]+}", handlers.UpdateUser)
	router.Delete("/users/{id:[0-9]+}", handlers.DeleteUser)

	router.Post("/products", handlers.AddProduct)
	router.Get("/products", handlers.ProductList)
	router.Post("/products/{id:[0-9]+}", handlers.UpdateProduct)
	router.Delete("/products/{id:[0-9]+}", handlers.DeleteProduct)

	router.Post("/orders", handlers.AddOrder)
	router.Get("/orders", handlers.OrderList)
	router.Post("/orders/{id:[0-9]+}", handlers.UpdateOrder)
	router.Delete("/orders/{id:[0-9]+}", handlers.DeleteOrder)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}
	srv.Addr = fmt.Sprintf("%s:%s", url, port)
	err := srv.ListenAndServe()
	fmt.Printf("Server is stoped...: %s\n", err.Error())
}
