package main

import (
	"log"
	"net/http"

	"learn-go/internal/db"
	"learn-go/internal/router"
)

func main() {
    db.InitDB()
    r := router.SetupRouter()

    log.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
