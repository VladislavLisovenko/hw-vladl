package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func encodedMessage(text string) ([]byte, error) {
	message := &Message{
		Content: text,
	}

	return json.Marshal(message)
}

func Test_handler(t *testing.T) {
	encodedMessage, err := encodedMessage("test")
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/health-check", bytes.NewReader(encodedMessage))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"content":"Message 'test', got"}`
	fmt.Println(rr.Body.String())
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
