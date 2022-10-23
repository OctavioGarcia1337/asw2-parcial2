package model

type OrderDetail struct {
	OrderDetailId int     `gorm:"primaryKey;AUTO_INCREMENT"`
	OrderId       int     `gorm:"type:int"`
	ProductId     int     `gorm:"type:int"`
	Quantity      int     `gorm:"type:int;not null;"`
	Price         float32 `gorm:"type:decimal;not null;"`
	CurrencyId    string  `gorm:"type:varchar(10);not null;"`
	Name          string  `gorm:"type:varchar(100);not null;"`
}

type OrderDetails []OrderDetail
