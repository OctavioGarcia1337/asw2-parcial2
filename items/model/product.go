package model

type Product struct {
	ProductId   int     `gorm:"primaryKey;AUTO_INCREMENT"`
	CategoryId  int     `gorm:"foreignKey:CategoryId"`
	Name        string  `gorm:"type:varchar(100);not null;unique;"`
	Description string  `gorm:"type:varchar(255);not null;"`
	Price       float32 `gorm:"type:decimal;not null;"`
	CurrencyId  string  `gorm:"type:varchar(10);not null;"`
	Stock       int     `gorm:"type:int;not null;"`
	Picture     string  `gorm:"type:varchar(100);not null;"`
}

type Products []Product
