package main

import (
	"learn-go/internal/router"
	"log"
	"net/http"
)

func main() {
    r := router.SetupRouter()
    log.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
