package main

import (
	"fmt"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/repositories/message_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/room_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"log"
	"time"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := configs.NewDB(cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	//userRepo := user_repo.NewUserRepo(db)
	messageRepo := message_repo.NewMessageRepo(db)
	roomRepo := room_repo.NewRoomRepo(db)

	if err := db.AutoMigrate(user_repo.User{}); err != nil {
		log.Println(err)
		return
	}
	if err := db.AutoMigrate(room_repo.Room{}); err != nil {
		log.Println(err)
		return
	}
	if err := db.AutoMigrate(message_repo.Message{}); err != nil {
		log.Println(err)
		return
	}
	r := room_repo.Room{
		Name:    "first_chat",
		OwnerID: 1,
		//Type:      1,
		CreatedAt: time.Now(),
	}
	m := message_repo.Message{
		Text:      "hello guys",
		OwnerID:   1,
		RoomID:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if _, err := roomRepo.CreateRoom(r); err != nil {
		log.Println(fmt.Errorf("error while creating room:%v", err))
		return
	}
	if _, err := messageRepo.CreateMessage(m); err != nil {
		log.Println(fmt.Errorf("error while creating room:%v", err))
		return
	}
}
