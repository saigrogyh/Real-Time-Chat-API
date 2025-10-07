package repository

import ("gorm.io/gorm"
		. "github.com/saigrogyh/Real-Time-Chat-API/internal/domain"
)

type ChatRepository interface {
	Create(chat *Chat) error
	FindById(id uint) (*Chat, error)
	FindAll() ([]Chat ,error)
	Delete(id uint) error
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *chatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) Create(chat *Chat) error{
	return r.db.Create(chat).Error
}

func (r *chatRepository) FindById(id uint) (*Chat, error) {
	var chat Chat
	if err := r.db.Where("id = ?",id).First(&chat).Error; err != nil{
		return nil, err
	}
	return &chat,nil
}

func (r *chatRepository) FindAll() ([]Chat, error) {
	var chats []Chat
	if err := r.db.Find(&chats).Error; err != nil {
		return nil, err
	}
	return chats, nil
}

func (r *chatRepository) Delete(id uint) error {
	return r.db.Delete(&User{},id).Error
}
