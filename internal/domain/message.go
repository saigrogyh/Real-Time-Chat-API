package domain

import "time"


type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ChatID    uint      `gorm:"index" json:"chat_id"`
	SenderID  uint      `gorm:"index" json:"sender_id"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
