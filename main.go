package main

import (
	"log"

	"github.com/4rufi/curso-go/db"
	"github.com/4rufi/curso-go/handlers"
)
func main()  {
	if !db.ConnectionCheck() {
		log.Fatal("Error connect DB")
	}
	handlers.RouteHandlers()
}