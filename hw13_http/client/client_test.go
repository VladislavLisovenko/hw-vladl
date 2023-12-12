package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_encodedMessage(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				text: "Message for server",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := encodedMessage(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodedMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_sendRequest(t *testing.T) {
	type args struct {
		url     string
		port    string
		method  string
		message []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should be error, because the server is unavailable",
			args: args{
				url:     "http://localhost",
				port:    "8080",
				method:  "GET",
				message: []byte("sdfsdf"),
			},
			wantErr: true,
		},
		{
			name: "Should be no error, because the server is available",
			args: args{
				url:     "https://go.dev/",
				port:    "80",
				method:  "GET",
				message: []byte("sdfsdf"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := sendRequest(tt.args.url, tt.args.port, tt.args.method, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_handleResponse(t *testing.T) {
	type args struct {
		url     string
		port    string
		method  string
		message []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Response should be nil, because the server is unavailable",
			args: args{
				url:     "http://localhost",
				port:    "8080",
				method:  "GET",
				message: []byte("sdfsdf"),
			},
			want: true,
		},
		{
			name: "Response should be non empty, because the server is available",
			args: args{
				url:     "https://go.dev/",
				port:    "80",
				method:  "GET",
				message: []byte("sdfsdf"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, _ := sendRequest(tt.args.url, tt.args.port, tt.args.method, tt.args.message)
			got := response == nil
			require.Equal(t, got, tt.want)
		})
	}
}
