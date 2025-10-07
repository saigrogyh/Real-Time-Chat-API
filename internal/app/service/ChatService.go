package service

import (
	// "log"
	// . "github.com/saigrogyh/Real-Time-Chat-API/internal/app"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app/repository"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/domain"
	"gorm.io/gorm"
)

type ChatService struct{
	repo ChatRepository
}

func CreateChat(db *gorm.DB, chat *Chat) error {
	result := db.Create(chat)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetChat(db *gorm.DB,id uint) (*Chat,error) {
	var chat Chat
	result := db.First(&chat, id)
	if result.Error != nil {
		return nil,result.Error

	}

	return &chat , nil
}

func UpdateChat(db *gorm.DB, chat *Chat,id uint) error {
	_, err := GetChat(db, id)
	if err != nil {
		return err // ถ้าไม่เจอ chat ก็รีเทิร์นเลย
	}

	result := db.Save(&chat)

	if result.Error != nil {
		return result.Error
	}

	return nil
}