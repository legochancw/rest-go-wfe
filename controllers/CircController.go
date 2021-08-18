package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-go-wfe/entity"

	"github.com/gorilla/mux"
)

//GetAllTmplCirc get all Circulation
func (app *App) GetAllCirc(w http.ResponseWriter, r *http.Request) {
	var circ_list []entity.CIRC

	rows, err := app.Database.Query("SELECT * FROM CIRC")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var circ entity.CIRC
		if err = rows.Scan(&circ.CIRC_ID,
			&circ.CIRC_TYPE,
			&circ.CIRC_TTL,
			&circ.CIRC_CAT,
			&circ.CIRC_SUB_CAT,
			&circ.STS,
			&circ.PROJ_ID,
			&circ.DUR,
			&circ.CRTD_TIME,
			&circ.UPTD_TIME); err != nil {
			fmt.Println("Scanning failed.....")
			fmt.Println(err.Error())
			return
		}
		if err != nil {
			log.Fatal(err)
		} else {

			circ_list = append(circ_list, circ)
		}
		defer rows.Close()
	}
	jsonResult, err := json.Marshal(circ_list)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResult))
}

//GetTmplCircByID returns Template Circulation with specific ID
func (app *App) GetCircByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	circ_id := vars["circ_id"]
	//log.Println("[GetTmplCircByID] tmpl_circ_id:" + tmpl_circ_id)

	var circ entity.CIRC
	err := app.Database.QueryRow("SELECT * FROM CIRC WHERE CIRC_ID = ?", circ_id).Scan(&circ.CIRC_ID,
		&circ.CIRC_TYPE,
		&circ.CIRC_TTL,
		&circ.CIRC_CAT,
		&circ.CIRC_SUB_CAT,
		&circ.STS,
		&circ.PROJ_ID,
		&circ.DUR)
	if err != nil {
		panic(err.Error())
	}
	jsonResult, err := json.Marshal(&circ)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResult))
}
