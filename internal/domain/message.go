package domain

import "time"

type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ChatID    uint      `gorm:"index;not null" json:"chat_id"`
	SenderID  uint      `gorm:"index;not null" json:"sender_id"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `json:"created_at"`

	Chat   Chat `gorm:"foreignKey:ChatID"`
	Sender User `gorm:"foreignKey:SenderID"`
}
