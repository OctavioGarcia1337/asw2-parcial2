package model

type User struct {
	UserId    int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Username  string `gorm:"type:varchar(40);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	Addresses Addresses
}

type Users []User
