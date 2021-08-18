package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-go-wfe/entity"

	"github.com/gorilla/mux"
)

//CIRC_TASK_ID
//CIRC_ID
//CIRC_TASK_TYPE
//CIRC_TASK_TTL
//CIRC_TASK_REM

//GetAllTmplCircTask get Tasks in the template circulation by CIRC_ID
func (app *App) GetTmplTaskByCircId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	circ_id := vars["circ_id"]
	var tmpl_circ_task_list []entity.TMPL_CIRC_TASK

	rows, err := app.Database.Query("SELECT * FROM TMPL_CIRC_TASK WHERE CIRC_ID = ?", circ_id)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var tmpl_circ_task entity.TMPL_CIRC_TASK
		if err = rows.Scan(&tmpl_circ_task.CIRC_TASK_ID,
			&tmpl_circ_task.CIRC_ID,
			&tmpl_circ_task.CIRC_TASK_TYPE,
			&tmpl_circ_task.CIRC_TASK_TTL,
			&tmpl_circ_task.CIRC_TASK_REM,
		); err != nil {
			fmt.Println("Scanning failed.....")
			fmt.Println(err.Error())
			return
		}
		if err != nil {
			log.Fatal(err)
		} else {
			tmpl_circ_task_list = append(tmpl_circ_task_list, tmpl_circ_task)
		}
		defer rows.Close()
	}
	jsonResult, err := json.Marshal(tmpl_circ_task_list)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResult))
}

//GetTmplTaskByTaskId returns Template Tasks by Task ID
func (app *App) GetTmplTaskByTaskId(w http.ResponseWriter, r *http.Request) {
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
