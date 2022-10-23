package model

type Order struct {
	ID         int     `gorm:"primaryKey;AUTO_INCREMENT"`
	UserId     int     `gorm:"type:int;not null"`
	AddressId  int     `gorm:"type:int;not null"`
	Date       string  `gorm:"type:date;not null;"`
	Total      float32 `gorm:"type:decimal;not null;"`
	CurrencyId string  `gorm:"type:varchar(10);not null;"`
}

type Orders []Order
