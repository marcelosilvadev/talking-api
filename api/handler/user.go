package handler

import (
	"database/sql"
	"encoding/json"
	"hackaton-facef-api/api/db"
	"hackaton-facef-api/model"
	"hackaton-facef-api/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//InsertUser ...
func InsertUser(w http.ResponseWriter, r *http.Request) {

	var t util.App
	var si model.User
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&si); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = si.InsertUser(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Usuário", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, si, 0, 0)
}

//GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {
	var si model.User
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetUser] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	si.Email = r.FormValue("email")
	err = si.GetUser(db)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetUser -  Não há Usuário com este Email.")
			t.ResponseWithError(w, 404, "Não há Usuário com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetUser -  Erro ao tentar buscar Usuário. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, si, 0, 0)
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {

	var t util.App
	var si model.User
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados está down", "")
		return
	}
	db := d.DB
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&si); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = si.Login(db)
	if err != nil {
		t.ResponseWithError(w, 404, "Usuário ou senha incorreta! Tente novamente.", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, si, 0, 0)
}

//GetRanking ...
func GetRanking(w http.ResponseWriter, r *http.Request) {
	var si model.User
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetRanking] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	res, err := si.GetRanking(db)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetRanking -  Não há Ranking.")
			t.ResponseWithError(w, 404, "Não há Ranking.", err.Error())
		} else {
			log.Printf("[handler/GetRanking -  Erro ao tentar buscar Ranking. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, res, 0, 0)
}

//UpdatePoints ...
func UpdatePoints(w http.ResponseWriter, r *http.Request) {
	var a model.User
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdatePoints] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	point, err := strconv.Atoi(vars["point"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	a.Id = int64(id)
	a.Points = strconv.Itoa(point)
	if err := a.UpdatePoints(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, a, 0, 0)
}
