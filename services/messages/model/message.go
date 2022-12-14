package model

type Message struct {
	ID        int    `gorm:"primaryKey;AUTO_INCREMENT"`
	UserId    int    `gorm:"type:int;not null"`
	ItemId    string `gorm:"type:varchar(255);not null"`
	Body      string `gorm:"type:varchar(255);not null"`
	CreatedAt string `gorm:"type:datetime"`
	System    bool   `gorm:"type:bool""`
}

type Messages []Message
