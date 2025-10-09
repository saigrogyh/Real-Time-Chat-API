package service

import (
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/domain"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app/repository"
)

type ChatService struct {
	repo ChatRepository
}

func NewChatService(repo ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) CreateChat(title string) (*Chat, error) {
	chat := &Chat{
		Title: title,
	}
	err := s.repo.Create(chat)
	return chat, err
}

func (s *ChatService) GetChatByID(id uint) (*Chat, error) {
	return s.repo.FindById(id)
}

// func (s *ChatService) GetChatWithMessages(id uint) (*Chat, error) {
// 	return s.repo.FindByIDWithMessages(id)
// }


func (s *ChatService) UpdateChat(chat *Chat) error {
	_, err := s.repo.FindById(chat.ID)
	if err != nil {
		return err
	}
	return s.repo.Update(chat)
}


func (s *ChatService) DeleteChat(id uint) error {
	return s.repo.Delete(id)
}
