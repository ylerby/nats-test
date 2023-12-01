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

	/* model := ModelJson{
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
	err = sc.Publish("nats-channel", data) */

	model2 := ModelJson{
		OrderUid:    "31",
		TrackNumber: "iuuuu",
		Entry:       "bbeqf",
		Locale:      "efq",
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
			Name:    "your name is qedcqw",
			Phone:   "123-456-7890784421e12e13",
			Zip:     "1234542r34",
			City:    "Citereqrdeqy",
			Address: "123 Stree54223rt",
			Region:  "Regiovvdsn",
			Email:   "john.doe@f3d132r31example.com",
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
			Transaction:  "ABCr3r31123",
			RequestId:    "R3131EQ-001",
			Currency:     "US1314133131D",
			Provider:     "Pa53513431yPal",
			Amount:       11110,
			PaymentDt:    1613200,
			Bank:         "XY343r3r3rZ Bank",
			DeliveryCost: 130,
			GoodsTotal:   9440,
			CustomFee:    53,
		},
	}

	// Создание и заполнение слайса Items
	items2 := []struct {
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
			ChrtId:      999,
			TrackNumber: "nuuuum",
			Price:       10000,
			Rid:         "riiiiiiiid",
			Name:        "Nameeee",
			Sale:        234,
			Size:        "xxxxxl",
			TotalPrice:  33333,
			NmId:        00001,
			Brand:       "Gucci",
			Status:      345,
		},
		{
			ChrtId:      13145,
			TrackNumber: "nuuuuuuuuuuuuum",
			Price:       42112,
			Rid:         "riiiiiiiiiiiiiid2",
			Name:        "itmee2",
			Sale:        3132,
			Size:        "xxxxxxxxxxxxxs",
			TotalPrice:  353,
			NmId:        105555,
			Brand:       "lv",
			Status:      313,
		},
	}

	model2.Items = items2

	data, err := json.Marshal(model2)
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
