package publisher

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
)

type ModelJson struct {
	OrderUid    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int    `json:"amount"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost int    `json:"delivery_cost"`
		GoodsTotal   int    `json:"goods_total"`
		CustomFee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int    `json:"total_price"`
		NmId        int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	} `json:"items"`
	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerId        string `json:"customer_id"`
	DeliveryService   string `json:"delivery_service"`
	Shardkey          string `json:"shardkey"`
	SmId              int    `json:"sm_id"`
	DateCreated       string `json:"date_created"`
	OofShard          string `json:"oof_shard"`
}

type Nats struct {
	ClusterId,
	ClientId string
}

type NatsInterface interface {
	PublicMessage()
	CirclePublicMessage()
}

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

	//todo: намокать тестовых значений

	model := ModelJson{
		OrderUid:    "12345",
		TrackNumber: "T123",
		Entry:       "Some entry",
		Locale:      "en",
		// Заполнение структуры для Delivery
		Delivery: struct {
			Name    string `json:"name"`
			Phone   string `json:"phone"`
			Zip     string `json:"zip"`
			City    string `json:"city"`
			Address string `json:"address"`
			Region  string `json:"region"`
			Email   string `json:"email"`
		}{
			Name:    "John Doe",
			Phone:   "123-456-7890",
			Zip:     "12345",
			City:    "City",
			Address: "123 Street",
			Region:  "Region",
			Email:   "john.doe@example.com",
		},
		// Заполнение структуры для Payment
		Payment: struct {
			Transaction  string `json:"transaction"`
			RequestId    string `json:"request_id"`
			Currency     string `json:"currency"`
			Provider     string `json:"provider"`
			Amount       int    `json:"amount"`
			PaymentDt    int    `json:"payment_dt"`
			Bank         string `json:"bank"`
			DeliveryCost int    `json:"delivery_cost"`
			GoodsTotal   int    `json:"goods_total"`
			CustomFee    int    `json:"custom_fee"`
		}{
			Transaction:  "ABC123",
			RequestId:    "REQ-001",
			Currency:     "USD",
			Provider:     "PayPal",
			Amount:       100,
			PaymentDt:    1618253200,
			Bank:         "XYZ Bank",
			DeliveryCost: 10,
			GoodsTotal:   90,
			CustomFee:    5,
		},
	}

	// Создание и заполнение слайса Items
	items := []struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int    `json:"total_price"`
		NmId        int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	}{
		{
			ChrtId:      1,
			TrackNumber: "T123",
			Price:       50,
			Rid:         "RID001",
			Name:        "Item 1",
			Sale:        10,
			Size:        "M",
			TotalPrice:  45,
			NmId:        1001,
			Brand:       "Brand A",
			Status:      1,
		},
		{
			ChrtId:      2,
			TrackNumber: "T123",
			Price:       40,
			Rid:         "RID002",
			Name:        "Item 2",
			Sale:        5,
			Size:        "L",
			TotalPrice:  38,
			NmId:        1002,
			Brand:       "Brand B",
			Status:      1,
		},
	}

	model.Items = items
	data, err := json.Marshal(model)
	err = sc.Publish("nats-channel", data)

	model3 := ModelJson{
		OrderUid:    "1",
		TrackNumber: "trck",
		Entry:       "entry",
		Locale:      "loc",
		// Заполнение структуры для Delivery
		Delivery: struct {
			Name    string `json:"name"`
			Phone   string `json:"phone"`
			Zip     string `json:"zip"`
			City    string `json:"city"`
			Address string `json:"address"`
			Region  string `json:"region"`
			Email   string `json:"email"`
		}{
			Name:    "name...",
			Phone:   "phone",
			Zip:     "zip",
			City:    "city",
			Address: "street 13a",
			Region:  "Moscow",
			Email:   "test@mail.ru",
		},
		// Заполнение структуры для Payment
		Payment: struct {
			Transaction  string `json:"transaction"`
			RequestId    string `json:"request_id"`
			Currency     string `json:"currency"`
			Provider     string `json:"provider"`
			Amount       int    `json:"amount"`
			PaymentDt    int    `json:"payment_dt"`
			Bank         string `json:"bank"`
			DeliveryCost int    `json:"delivery_cost"`
			GoodsTotal   int    `json:"goods_total"`
			CustomFee    int    `json:"custom_fee"`
		}{
			Transaction:  "transaction",
			RequestId:    "req_id",
			Currency:     "currency",
			Provider:     "nahabinoRU",
			Amount:       32,
			PaymentDt:    12,
			Bank:         "eee bank",
			DeliveryCost: 123,
			GoodsTotal:   94446,
			CustomFee:    535,
		},
	}

	// Создание и заполнение слайса Items
	items3 := []struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int    `json:"total_price"`
		NmId        int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	}{
		{
			ChrtId:      334,
			TrackNumber: "trcknum",
			Price:       1055,
			Rid:         "ridddd",
			Name:        "nm",
			Sale:        2312,
			Size:        "0",
			TotalPrice:  31,
			NmId:        38,
			Brand:       "brnd",
			Status:      89,
		},
		{
			ChrtId:      1318,
			TrackNumber: "tracccknuuuuuuuuuuuuum",
			Price:       421152,
			Rid:         "ridddd2",
			Name:        "itms2",
			Sale:        313,
			Size:        "000",
			TotalPrice:  3548,
			NmId:        10558,
			Brand:       "unknown",
			Status:      404,
		},
	}

	model3.Items = items3

	data, err = json.Marshal(model3)
	err = sc.Publish("nats-channel", data)
	//n.CirclePublicMessage()

	defer func() {
		err = sc.Close()
		if err != nil {
			log.Printf("ошибка при закрытии подключения %s", err)
		}
	}()
}

func (n *Nats) CirclePublicMessage() {
	for {
		break
	}
}
