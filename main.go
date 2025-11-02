package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/saigrogyh/Real-Time-Chat-API/internal/app/auth"
	"github.com/saigrogyh/Real-Time-Chat-API/internal/app/handler"
	"github.com/saigrogyh/Real-Time-Chat-API/internal/app/repository"
	"github.com/saigrogyh/Real-Time-Chat-API/internal/app/service"
	"github.com/saigrogyh/Real-Time-Chat-API/internal/domain"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Fatal("JWT_SECRET is not set in .env file")
    }

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("CHATDB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	fmt.Println("Environment variables loaded successfully")
	fmt.Println("DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database:")
	}

	fmt.Println("Database connection established successfully")

	err = db.AutoMigrate(&domain.User{}, &domain.Message{}, &domain.Chat{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	chatRepo := repository.NewChatRepository(db)
	messageRepo := repository.NewMessageRepository(db)

	userService := service.NewUserService(userRepo)
	chatService := service.NewChatService(chatRepo)
	messageService := service.NewMessageService(messageRepo)

	userHandler := handler.NewUserHandler(*userService)
	chatHandler := handler.NewChatHandler(*chatService)
	messageHandler := handler.NewMessageHandler(*messageService)

	app := fiber.New()

	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)

	// Group protected routes with JWT middleware
	api := app.Group("/api", auth.JWTProtected(jwtSecret)) 

	// Chat routes
	api.Post("/chats", chatHandler.CreateChat)
	api.Get("/chats/:id", chatHandler.GetChatByChatID)
	api.Get("/chats/delete/:id", chatHandler.DeleteChat)

	// Message routes
	api.Post("/messages", messageHandler.SendMessage)
	api.Get("/chats/:id/messages", messageHandler.GetAllMsgFromChatID)
	api.Get("/users/:id/messages", messageHandler.GetMessagesByUserID)

	// Start server
	port := os.Getenv("APP_PORT")
	app.Listen(":" + port)

}
