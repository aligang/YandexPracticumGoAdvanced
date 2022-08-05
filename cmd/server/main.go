package main

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/handler"
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

	handler1 := handler.ApiHandler{}
	//handler2 := ApiHandler2{}

	http.DefaultServeMux.Handle("/update/", handler1)
	//http.DefaultServeMux.Handle("/update/gauge/BuckHashSys/", handler2)
	server.Serve(listener)

}
