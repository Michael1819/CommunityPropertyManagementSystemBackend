package main

import (
	"backend/backend"
	"backend/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is running on port 8080")

	backend.Init()
	defer backend.Close()

	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}
