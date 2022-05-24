package main

import (
	"fmt"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDb(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := NewDb(cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user_repo.NewUserRepo(db)
	if err := db.AutoMigrate(user_repo.User{}); err != nil {
		log.Println(err)
		return
	}
	//db.AutoMigrate(message_repo.Message{})
	//db.AutoMigrate(room_repo.Room{})

	//u := user_repo.User{
	//	Name:      "John",
	//	CreatedAt: time.Now(),
	//}

	//user, err := userRepo.CreateUser(u)
	user, err := userRepo.GetUser(1)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(user.UpdatedAt)
	// Tests
}
