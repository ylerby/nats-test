package pub

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/nats-io/stan.go"
	"log"
	"sync"
)

func New(cluster, client string) NatsInterface {
	return &Nats{
		ClusterId: cluster,
		ClientId:  client,
	}
}

func (n *Nats) PublicMessage() {
	log.Printf("%s, %s", n.ClientId, n.ClusterId)
	sc, err := stan.Connect(n.ClusterId, n.ClientId)
	if err != nil {
		log.Printf("ошибка при подключении %s", err)
	}

	models := make([]ModelJson, mockSize)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for i := 0; i < mockSize; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			mu.Lock()
			var model ModelJson
			err = gofakeit.Struct(&model)
			if err != nil {
				log.Printf("ошибка при создании mock структуры")
			}
			models[index] = model
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	for i := 0; i < mockSize; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			mu.Lock()
			data, err := json.Marshal(models[index])
			if err != nil {
				log.Printf("ошибка при сериализации json-файла %s", err)
			}

			err = sc.Publish("nats-channel", data)

			if err != nil {
				log.Printf("ошибка при добавлении файла %s", err)
			}

			mu.Unlock()
		}(i)
	}

	wg.Wait()

	err = sc.Publish("nats-channel", []byte("incorrect value"))
	if err != nil {
		log.Printf("ошибка при добавлении файла %s", err)
	}

	defer func() {
		err = sc.Close()
		if err != nil {
			log.Printf("ошибка при закрытии подключения %s", err)
		}
	}()
}
