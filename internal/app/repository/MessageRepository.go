package repository

import (
	"gorm.io/gorm"
	"github.com/saigrogyh/Real-Time-Chat-API/internal/domain"
)

type MessageRepository interface {
	Create(message *domain.Message) error
	GetAll() ([]domain.Message, error)
	GetByID(id uint) (*domain.Message, error)
	GetByChatID(chatID uint) ([]domain.Message, error)
	GetByUserID(userID uint) ([]domain.Message, error)
	Delete(id uint) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(message *domain.Message) error {
	return r.db.Create(message).Error
}

func (r *messageRepository) GetAll() ([]domain.Message, error) {
	var messages []domain.Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) GetByID(id uint) (*domain.Message, error) {
	var message domain.Message
	err := r.db.First(&message, id).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *messageRepository) GetByChatID(chatID uint) ([]domain.Message, error) {
	var messages []domain.Message
	err := r.db.Where("chat_id = ?", chatID).Find(&messages).Error
	return messages, err
}

func (r *messageRepository) GetByUserID(userID uint) ([]domain.Message, error) {
	var messages []domain.Message
	err := r.db.Where("sender_id = ?", userID).Find(&messages).Error
	return messages, err
}

func (r *messageRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Message{}, id).Error
}