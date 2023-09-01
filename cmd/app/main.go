package main

import (
	"avito-backend-task/internal/controller/http/v1"
	"avito-backend-task/internal/entity"
	"avito-backend-task/internal/repo/pgdb"
	"avito-backend-task/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
)

func MigrateDB(db *gorm.DB) {
	err := entity.MigrateSegments(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}
	err = entity.MigrateUsers(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}
	err = entity.MigrateUserSegmentPairs(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}
	MigrateDB(db)
	segRepo := &pgdb.SegmentRepo{DB: db}
	userRepo := &pgdb.UserRepo{DB: db}
	app := fiber.New()
	v1.NewRouter(app, v1.Repos{SegmentRepo: segRepo, UserRepo: userRepo})
	app.Listen(":8080")
}
