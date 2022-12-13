package model

type User struct {
	ID        int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Username  string `gorm:"type:varchar(45);not null;unique"`
	Password  string `gorm:"type:varchar(45);not null"`
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
}

type Users []User
