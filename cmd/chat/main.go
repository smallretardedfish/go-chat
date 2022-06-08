package main

import (
	"fmt"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/domains/chat"
	"github.com/smallretardedfish/go-chat/internal/domains/user"
	"github.com/smallretardedfish/go-chat/internal/repositories/message_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/room_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_cred_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
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

	userRepo := user_repo.NewUserRepo(db)
	roomRepo := room_repo.NewRoomRepo(db)
	messageRepo := message_repo.NewMessageRepo(db)
	credsRepo := user_cred_repo.NewUserCredentialsRepo(db)

	messageSvc := chat.NewMessageServiceImpl(messageRepo)
	roomSvc := chat.NewRoomServiceImpl(roomRepo)
	authSvc := user.NewAuthServiceImpl(credsRepo, userRepo)
	//userSvc := user.NewUserServiceImpl(userRepo)

	// testing here
	newUser, err := authSvc.SignUp(user.User{
		Name:  "Victor32284",
		Email: "victor32284@mail.ua",
	}, user.UserCredentials{
		Email:    "victor32284@mail.ua",
		Password: "victorPASS",
	})

	newRoom, err := roomSvc.CreateRoom(chat.Room{
		Name:    "NEW ROOM",
		OwnerID: newUser.ID,
	}, nil)

	room, err := roomSvc.GetRoom(newUser.ID, newRoom.ID)
	if err != nil {
		return
	}
	fmt.Println("Owner Name is", room.OwnerID, room.Owner.Name, room.Users)
	_, err = messageSvc.CreateMessage(chat.Message{
		Text:    "HELLO",
		OwnerID: newUser.ID,
		RoomID:  room.ID,
	})
	if err != nil {
		log.Println(err)
		return
	}

	messages, err := messageSvc.GetMessages(nil, nil, newUser.ID, room.ID)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(messages)
	fmt.Println("message is written by: ", messages[0].Owner.Name, messages[0].Owner.ID)

	return
}
