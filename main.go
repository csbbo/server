package main

import (
	"log"
	"net/http"
	_ "server/server"
)

func main() {
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
