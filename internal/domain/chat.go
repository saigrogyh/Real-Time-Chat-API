package domain

import "time"

type Chat struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"size:255" json:"title"`
	CreatedAt time.Time `json:"created_at"`

	Messages []Message `gorm:"foreignKey:ChatID"` // relation กับข้อความในห้อง
}
