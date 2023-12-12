package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
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

func encodedMessage(text string) ([]byte, error) {
	message := &Message{
		Content: text,
		Date:    time.Now().Format(time.DateOnly),
	}

	return json.Marshal(message)
}

func sendRequest(url string, port string, method string, message []byte) (*http.Response, error) {
	getRequest, err := http.NewRequest(method, fmt.Sprintf("%s:%s", url, port), bytes.NewReader(message))
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

func handleResponse(response *http.Response) error {
	var responseBody Response
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return err
	}
	fmt.Println(responseBody.Content)
	return nil
}

func main() {
	url := ""
	port := ""
	text := ""
	method := ""
	flag.StringVar(&url, "url", "http://localhost", "URL to send to")
	flag.StringVar(&port, "port", "8080", "Port to send to")
	flag.StringVar(&text, "text", "", "Message text to send")
	flag.StringVar(&method, "method", http.MethodGet, "Method to send")
	flag.Parse()

	encodedMessage, err := encodedMessage(text)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	response, err := sendRequest(url, port, method, encodedMessage)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	err = handleResponse(response)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
