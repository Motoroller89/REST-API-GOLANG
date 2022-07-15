package main

import (
	"awesomeProject1/internal/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	router := httprouter.New()

	handler := user.NewHandler()

	handler.Register(router)

	start(router)

}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))
}
