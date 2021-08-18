package main

import (
	"log"
	"net/http"
	"rest-go-wfe/controllers"
	"rest-go-wfe/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db := database.InitConn()
	log.Println("Starting the HTTP server on port 8090")

	app := &controllers.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: db,
	}
	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8090", app.Router))
}
