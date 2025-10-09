package service

import (
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app/repository"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/domain"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app"
)


type MessageService struct {
	repo MessageRepository
}

func NewMessageService(r MessageRepository) *MessageService {
	return &MessageService{repo: r}
}

func (s *MessageService) SendMessage(message *SendMessageRequest) (Message, error) {
	msg := &Message{
		ChatID:   message.ChatID,
		SenderID: message.UserID,
		Content:  message.Content,
	}

	err := s.repo.Create(msg)
	if err != nil { 
		return Message{}, err
	}
	return *msg, nil
}

func (s *MessageService) GetMessagesByChatID(chatID uint) ([]Message, error) {
	return s.repo.GetByChatID(chatID)
}

func (s *MessageService) GetMessagesByUserID(userID uint) ([]Message, error) {
	return s.repo.GetByUserID(userID)
}

func (s *MessageService) DeleteMessage(id uint) error {
	return s.repo.Delete(id)
}