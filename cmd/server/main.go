package main

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/handler"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage"
	"net"
	"net/http"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server := &http.Server{}

	update_handler := handler.ApiHandler{
		Storage: storage.New(),
	}
	http.DefaultServeMux.Handle("/update/", update_handler)
	server.Serve(listener)

}
