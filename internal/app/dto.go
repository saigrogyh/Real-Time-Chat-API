package app

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SendMessageRequest struct {
	ChatID  uint   `json:"chat_id" validate:"required"`
	UserID  uint   `json:"user_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

