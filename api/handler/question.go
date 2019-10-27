package handler

import (
	"database/sql"
	"hackaton-facef-api/api/db"
	"hackaton-facef-api/model"
	"hackaton-facef-api/util"
	"log"
	"net/http"
)

//GetQuestion ...
func GetQuestion(w http.ResponseWriter, r *http.Request) {
	var si model.Questions
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetQuestion] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	res, err := si.GetQuestions(db)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetQuestion -  Não há Questão.")
			t.ResponseWithError(w, 404, "Não há Usuário com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetQuestion -  Erro ao tentar buscar Questão. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, res, 0, 0)
}
