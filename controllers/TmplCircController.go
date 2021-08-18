package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-go-wfe/entity"

	"github.com/gorilla/mux"
)

//GetAllTmplCirc get all template Circulation
func (app *App) GetAllTmplCirc(w http.ResponseWriter, r *http.Request) {
	var tmpl_circ_list []entity.TMPL_CIRC

	rows, err := app.Database.Query("SELECT * FROM TMPL_CIRC")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var tmpl_circ entity.TMPL_CIRC
		if err = rows.Scan(&tmpl_circ.CIRC_ID,
			&tmpl_circ.CIRC_TYPE,
			&tmpl_circ.CIRC_TTL,
			&tmpl_circ.CIRC_CAT,
			&tmpl_circ.CIRC_SUB_CAT,
			&tmpl_circ.STS,
			&tmpl_circ.PROJ_ID,
			&tmpl_circ.DUR,
		); err != nil {
			fmt.Println("Scanning failed.....")
			fmt.Println(err.Error())
			return
		}
		if err != nil {
			log.Fatal(err)
		} else {
			tmpl_circ_list = append(tmpl_circ_list, tmpl_circ)
		}
		defer rows.Close()
	}
	jsonResult, err := json.Marshal(tmpl_circ_list)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResult))
}

//GetTmplCircByID returns Template Circulation with specific ID
func (app *App) GetTmplCircByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tmpl_circ_id := vars["tmpl_circ_id"]
	//log.Println("[GetTmplCircByID] tmpl_circ_id:" + tmpl_circ_id)

	var tmpl_circ entity.TMPL_CIRC
	err := app.Database.QueryRow("SELECT * FROM TMPL_CIRC WHERE CIRC_ID = ?", tmpl_circ_id).Scan(&tmpl_circ.CIRC_ID,
		&tmpl_circ.CIRC_TYPE,
		&tmpl_circ.CIRC_TTL,
		&tmpl_circ.CIRC_CAT,
		&tmpl_circ.CIRC_SUB_CAT,
		&tmpl_circ.STS,
		&tmpl_circ.PROJ_ID,
		&tmpl_circ.DUR)
	if err != nil {
		panic(err.Error())
	}
	jsonResult, err := json.Marshal(&tmpl_circ)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResult))
}
