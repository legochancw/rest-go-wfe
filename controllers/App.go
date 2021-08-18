package controllers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	//router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	app.Router.HandleFunc("/tmpl/getCirc", app.GetAllTmplCirc).Methods("GET")
	app.Router.HandleFunc("/tmpl/getCirc/{tmpl_circ_id}", app.GetTmplCircByID).Methods("GET")

	app.Router.HandleFunc("/circ/get", app.GetAllCirc).Methods("GET")
	app.Router.HandleFunc("/circ/get/{circ_id}", app.GetCircByID).Methods("GET")

	//outer.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	//router.HandleFunc("/delete/{id}", controllers.DeletPersonByID).Methods("DELETE")
}
