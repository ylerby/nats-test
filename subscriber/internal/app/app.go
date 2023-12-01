package app

import (
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	jsonModel "subscriber/api/json"
	"subscriber/internal/database/cache"
	"subscriber/internal/database/sql"
	"subscriber/internal/service"
	"time"
)

type App struct {
	Server *http.Server
	Cache  cache.InterfaceCache
	Sql    sql.InterfaceSql
	Nats   service.InterfaceNats
}

func New() *App {
	return &App{
		Server: &http.Server{Addr: ":8080"},
		Cache:  cache.New(),
		Sql:    sql.New(),
		Nats:   service.New("test-cluster", "client-1", "nats-channel"),
	}
}

//todo: 1) nats-streaming в docker-compose up
//todo: 2) МОДЕЛИ БЛЯТЬ !!!!!!!!!!!!!!!!!!!!!

func (a *App) Run() {
	log.Println("запуск")
	a.Cache.Connect()
	log.Println("подключен кеш")
	sc, channel, err := a.Nats.Connect()
	if err != nil {
		log.Printf("ошибка при работе с nats %s", err)
	}

	log.Println("подключен натс")

	err = a.Sql.Connect()

	log.Println("подключен sql")

	if err != nil {
		log.Printf("ошибка при подключении к БД %s", err)
	}

	sub, err := sc.Subscribe(channel, a.Sub, stan.DeliverAllAvailable())
	if err != nil {
		log.Printf("ошибка при подписке %s", err)
	}

	log.Println("подключен канал")

	defer func() {
		err = sub.Close()
		if err != nil {
			log.Printf("ошибка при закрытии %s", err)
		}
	}()

	res := a.Sql.GetAllRecords()
	log.Println("получены значения")

	log.Println(res)
	a.Cache.CacheDownloading(res)

	http.HandleFunc("/", a.Get)

	log.Println("запуск сервера")

	err = a.Server.ListenAndServe()
	if err != nil {
		log.Printf("ошибка при запуске сервера %s", err)
	}
}

func (a *App) Stop() {
	log.Println("завершение работы сервера")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		log.Fatalf("ошибка при завершении работы сервера %s", err)
	}
}

func (a *App) Sub(msg *stan.Msg) {
	model := jsonModel.ModelJson{}
	err := json.Unmarshal(msg.Data, &model)
	if err != nil {
		log.Fatalf("ошибка при десериализации %s", err)
	}

	//todo: добавить обработку ошибки при добавлении в кеш, sql-таблицу
	a.Cache.AddRecord(model)
	a.Sql.AddRecord(model)
}
