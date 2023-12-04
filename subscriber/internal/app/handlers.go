package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (a *App) Get(w http.ResponseWriter, r *http.Request) {
	formId := r.FormValue("id")
	id, err := strconv.Atoi(formId)
	if err != nil {
		log.Printf("ошибка при конвертации id записи %s", err)
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}

	val, ok := a.Cache.GetById(id)
	if ok {
		log.Println("попытка получения значения из кеша")
		response, err := json.Marshal(*val)
		if err != nil {
			log.Println("ошибка при сериализации объекта")
			http.Redirect(w, r, "/", http.StatusBadRequest)
		}
		log.Println("значение получено из кеша")
		_, err = w.Write(response)
		if err != nil {
			log.Println("ошибка при ответе")
			http.Redirect(w, r, "/", http.StatusBadRequest)
		}
	} else {
		log.Println("попытка получения значения из sql")
		sqlVal, ok := a.Sql.GetById(id)
		if !ok {
			_, err = w.Write([]byte("Заказ с таким id не найден"))
			if err != nil {
				log.Println("ошибка при ответе")
				http.Redirect(w, r, "/", http.StatusBadRequest)
			}
		} else {
			response, err := json.Marshal(*sqlVal)
			if err != nil {
				log.Println("ошибка при сериализации объекта")
				http.Redirect(w, r, "/", http.StatusBadRequest)
			}
			_, err = w.Write(response)
		}
	}
}
