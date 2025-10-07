package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "github.com/gofiber/swagger"
	// _ "github.com/saigrogyh/fiber-test/docs"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/domain"
)

// var HandlerDefault = New()

func main() {
	//load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Environment variables loaded successfully")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database:")
	}

	fmt.Println("Database connection established successfully")

	fmt.Println(db)
	// Initialize Fiber app
	app := fiber.New()

	// app.Post("/register", register)
	// app.Get("/login", login)

	
	// Swagger documentation
	// app.Get("/swagger/*", swagger.HandlerDefault)

	// Start server
	port := os.Getenv("APP_PORT")
	app.Listen(":"+port)

	db.AutoMigrate(&Chat{}, &User{})
}
