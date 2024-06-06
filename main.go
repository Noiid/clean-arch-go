package main

import (
	"clean_arch/delivery/http"
	"clean_arch/usecase"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	// Initialize the database connection
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Connected to the database successfully")

	userUseCase := usecase.NewUserUseCase(db)
	http.NewUserHandler(app, userUseCase) // Assuming you want to pass the db to your handler

	app.Listen(":3000")
}

// initDB initializes the database connection using GORM
func initDB() (*gorm.DB, error) {
	dsn := "host=localhost user=your_user password=your_password dbname=your_dbname port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
