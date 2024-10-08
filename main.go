package main

import (
	"go-blog/common"
	"go-blog/router"
	"log"
	"net/http"
)

func init() {
	common.LoadTemplate()
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
