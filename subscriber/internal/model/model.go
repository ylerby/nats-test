package models

import (
	"gorm.io/gorm"
	"time"
)

type Delivery struct {
	gorm.Model
	DeliveryID int    `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"type:varchar(255);unique"`
	Phone      string `gorm:"type:varchar(255);unique"`
	Zip        string `gorm:"type:varchar(255);unique"`
	City       string `gorm:"type:varchar(255);unique"`
	Address    string `gorm:"type:varchar(255);unique"`
	Region     string `gorm:"type:varchar(255);unique"`
	Email      string `gorm:"type:varchar(255);unique"`
}

type Payment struct {
	gorm.Model
	PaymentID    int    `gorm:"primaryKey;autoIncrement"`
	Transaction  string `gorm:"type:varchar(255);unique"`
	RequestId    string `gorm:"type:varchar(255);unique"`
	Currency     string `gorm:"type:varchar(255);unique"`
	Provider     string `gorm:"type:varchar(255);unique"`
	Amount       int    `gorm:"type:int"`
	PaymentDt    int    `gorm:"type:int"`
	Bank         string `gorm:"type:varchar(255);unique"`
	DeliveryCost int    `gorm:"type:int"`
	GoodsTotal   int    `gorm:"type:int"`
	CustomFee    int    `gorm:"type:int"`
}

type Items struct {
	gorm.Model
	ID          int    `gorm:"primaryKey;autoIncrement"`
	ChrtId      int    `gorm:"type:int;unique"`
	TrackNumber string `gorm:"type:varchar(255);unique"`
	Price       int    `gorm:"type:int"`
	Rid         string `gorm:"type:varchar(255);unique"`
	Name        string `gorm:"type:varchar(255);unique"`
	Sale        int    `gorm:"type:int"`
	Size        string `gorm:"type:varchar(255);unique"`
	TotalPrice  int    `gorm:"type:int"`
	NmId        int    `gorm:"type:int"`
	Brand       string `gorm:"type:varchar(255);unique"`
	Status      int    `gorm:"type:int"`
}

type Model struct {
	gorm.Model
	ID                int       `gorm:"primaryKey;autoIncrement"`
	OrderUid          string    `gorm:"type:varchar(255);unique"`
	TrackNumber       string    `gorm:"type:varchar(255);unique"`
	Entry             string    `gorm:"type:varchar(255);unique"`
	Delivery          Delivery  `gorm:"foreignKey:DeliveryID"` // Связь с таблицей Delivery
	Payment           Payment   `gorm:"foreignKey:PaymentID"`  // Связь с таблицей Payment
	Items             []Items   `gorm:"foreignKey:ID"`
	Locale            string    `gorm:"type:varchar(255);unique"`
	InternalSignature string    `gorm:"type:varchar(255);unique"`
	CustomerId        string    `gorm:"type:varchar(255);unique"`
	DeliveryService   string    `gorm:"type:varchar(255);unique"`
	ShardKey          string    `gorm:"type:varchar(255);unique"`
	SmId              int       `gorm:"type:int"`
	DateCreated       time.Time `gorm:"type:timestamp"`
	OofShard          string    `gorm:"type:varchar(255);unique"`
}

/*type Delivery struct {
	ID      uint   `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	Name    string `gorm:"type:varchar(255)"`
	Phone   string `gorm:"type:varchar(255)"`
	Zip     string `gorm:"type:varchar(255)"`
	City    string `gorm:"type:varchar(255)"`
	Address string `gorm:"type:varchar(255)"`
	Region  string `gorm:"type:varchar(255)"`
	Email   string `gorm:"type:varchar(255)"`
}

type Payment struct {
	ID           uint   `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	Transaction  string `gorm:"type:varchar(255)"`
	RequestId    string `gorm:"type:varchar(255)"`
	Currency     string `gorm:"type:varchar(255)"`
	Provider     string `gorm:"type:varchar(255)"`
	Amount       int
	PaymentDt    int
	Bank         string `gorm:"type:varchar(255)"`
	DeliveryCost int
	GoodsTotal   int
	CustomFee    int
}

type Items struct {
	ID          uint `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	ChrtId      int
	TrackNumber string `gorm:"type:varchar(255)"`
	Price       int
	Rid         string `gorm:"type:varchar(255)"`
	Name        string `gorm:"type:varchar(255)"`
	Sale        int
	Size        string `gorm:"type:varchar(255)"`
	TotalPrice  int
	NmId        int
	Brand       string `gorm:"type:varchar(255)"`
	Status      int
}

type Model struct {
	ID                uint     `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	OrderUid          string   `gorm:"type:varchar(255)"`
	TrackNumber       string   `gorm:"type:varchar(255)"`
	Entry             string   `gorm:"type:varchar(255)"`
	Delivery          Delivery `gorm:"foreignKey:ID"`
	Payment           Payment  `gorm:"foreignKey:ID"`
	Items             []Items  `gorm:"foreignKey:ID"`
	Locale            string   `gorm:"type:varchar(255)"`
	InternalSignature string   `gorm:"type:varchar(255)"`
	CustomerId        string   `gorm:"type:varchar(255)"`
	DeliveryService   string   `gorm:"type:varchar(255)"`
	ShardKey          string   `gorm:"type:varchar(255)"`
	SmId              int
	DateCreated       time.Time
	OofShard          string `gorm:"type:varchar(255)"`
} */
