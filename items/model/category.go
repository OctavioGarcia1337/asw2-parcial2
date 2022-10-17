package model

type Category struct {
	CategoryId  int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(255)"`
}

type Categories []Category
