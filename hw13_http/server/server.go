package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Content string `json:"content"`
	Date    string `json:"date"`
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
	responseText := fmt.Sprintf("Message '%s' from %s, got", message.Content, message.Date)

	response := &Response{
		Content: responseText,
	}

	return json.Marshal(response)
}

func handler(w http.ResponseWriter, r *http.Request) {
	message, err := decodedMessage(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	encodedResponse, err := encodedResponse(message)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	_, err = w.Write(encodedResponse)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
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
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
