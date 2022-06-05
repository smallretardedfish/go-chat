package main

import (
	"fmt"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/repositories/message_repo"
	//"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
	"log"
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
	//roomRepo := room_repo.NewRoomRepo(db)
	messageRepo := message_repo.NewMessageRepo(db)
	msg, err := messageRepo.CreateMessage(message_repo.Message{
		Text:         "HELLO ",
		OwnerID:      1,
		RoomID:       17,
		DeletedUsers: []int64{111},
	})

	fmt.Println(msg.ID)
	msg.DeletedUsers = append(msg.DeletedUsers, 2)
	updatedMessage, err := messageRepo.UpdateMessage(*msg)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(updatedMessage.DeletedUsers)
	messages, err := messageRepo.GetMessages(nil, 1, 17)

	for _, message := range messages {
		fmt.Println(message.Text)
	}
	return
}
