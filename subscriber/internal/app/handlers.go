package app

import (
	"log"
	"net/http"
	"strconv"
)

func (a *App) Get(w http.ResponseWriter, r *http.Request) {
	formId := r.FormValue("id")
	res := a.Sql.GetAllRecords()
	log.Println("res = ", res)
	id, err := strconv.Atoi(formId)
	if err != nil {
		log.Printf("ошибка при конвертации id записи %s", err)
	}

	//todo: сделать получение значения по id!
	val, ok := a.Cache.GetById(id)
	if ok {
		log.Println(*val)
	} else {
		//a.Sql.
	}

	_, err = w.Write([]byte("привет как жизнь"))
	if err != nil {
		return
	}
}
