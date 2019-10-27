package handler

import (
	"database/sql"
	"hackaton-facef-api/api/db"
	"hackaton-facef-api/model"
	"hackaton-facef-api/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetHistoric ...
func GetHistoric(w http.ResponseWriter, r *http.Request) {
	var dm model.Historic
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetHistoric] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	dm.User.Id = int64(id)
	err = dm.GetHistoric(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetHistoric -  Não há Hitorico com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Não há Historico com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetHistoric -  Erro ao tentar buscar Historico. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}
