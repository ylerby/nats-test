package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (a *App) Get(w http.ResponseWriter, r *http.Request) {
	formId := r.FormValue("id")
	log.Println("form id = ", formId)
	id, err := strconv.Atoi(formId)
	log.Println("conv id = ", id)
	log.Println("IDDDDD = ", id)
	if err != nil {
		log.Printf("ошибка при конвертации id записи %s", err)
	}

	val, ok := a.Cache.GetById(id)
	// todo: сделать получение значения из кеша, доработать кеш, добавить docker для nats-streaming
	// todo:
	if ok {
		log.Println("попытка получения значения из кеша")
		response, err := json.Marshal(*val)
		if err != nil {
			log.Println("ошибка при сериализации объекта")
		}
		_, err = w.Write([]byte(response))
		if err != nil {
			log.Println("ошибка при ответе")
		}
	} else {
		log.Println("попытка получения значения из sql")
		sqlVal, ok := a.Sql.GetById(id)
		if !ok {
			_, err = w.Write([]byte("Заказ с таким id не найден"))
			if err != nil {
				log.Println("ошибка при ответе")
				// todo: сделать response с ошибкой
			}
		} else {
			response, err := json.Marshal(*sqlVal)
			if err != nil {
				log.Println("ошибка при сериализации объекта")
			}
			_, err = w.Write([]byte(response))
		}
	}
}
