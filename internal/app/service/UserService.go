package service

import (
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/domain"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Register(r RegisterRequest) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &User{
		Username: r.Username,
		Email:    r.Email,
		Password: string(hashed),
	}

	return s.repo.Create(user)
}

func (s *UserService) Login(r LoginRequest) (*User, error) {
	user, err := s.repo.FindByUsername(r.Email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password)); err != nil {
		return nil, err
	}
	return user, nil
}
