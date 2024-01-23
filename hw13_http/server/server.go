package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Content string `json:"content"`
}

type Response struct {
	Content string `json:"content"`
}

func decodedMessage(body io.ReadCloser) (Message, error) {
	var message Message
	err := json.NewDecoder(body).Decode(&message)
	if err != nil {
		return message, err
	}

	return message, nil
}

func encodedResponse(message Message) ([]byte, error) {
	responseText := fmt.Sprintf("Message '%s', got", message.Content)

	response := &Response{
		Content: responseText,
	}

	return json.Marshal(response)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Methods allowed: GET, POST"))
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

	message := Message{Content: "GET request received"}
	if r.Method == http.MethodPost {
		var err error
		message, err = decodedMessage(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println(err.Error())
			}
			return
		}
	}

	encodedResponse, err := encodedResponse(message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

	_, err = w.Write(encodedResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
}

func main() {
	url := ""
	port := ""
	flag.StringVar(&url, "url", "localhost", "URL to listen from, without protocol, e.g. 'localhost'")
	flag.StringVar(&port, "port", "8080", "Port to listen from")
	flag.Parse()

	http.HandleFunc("/", handler)
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
