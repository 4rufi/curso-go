package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/4rufi/curso-go/middlew"
	"github.com/4rufi/curso-go/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* RouteHandlers Set port and cors with handlers */
func RouteHandlers()  {
	router := mux.NewRouter()

	router.HandleFunc("/registry", middlew.CheckDB(routers.Registry)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handlersCors := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":" + PORT, handlersCors))

}