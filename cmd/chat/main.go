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
	newUser1, err := authSvc.SignUp(user.User{
		Name:  "Oliver",
		Email: "oli1@mail.ua",
	}, user.UserCredentials{
		Email:    "oli1@mail.ua",
		Password: "victorPASS",
	})
	newUser2, err := authSvc.SignUp(user.User{
		Name:  "Ray",
		Email: "ray1@mail.ua",
	}, user.UserCredentials{
		Email:    "ray1@mail.ua",
		Password: "victorPASS",
	})

	newRoom, err := roomSvc.CreateRoom(chat.Room{
		Name:    "CONVERSATION",
		OwnerID: newUser1.ID,
	}, []int64{newUser2.ID})

	room, err := roomSvc.GetRoom(newUser2.ID, newRoom.ID)
	if err != nil {
		return
	}
	_, err = messageSvc.CreateMessage(chat.Message{
		Text:    "HELLO " + newUser2.Name,
		OwnerID: newUser1.ID,
		RoomID:  room.ID,
	})
	if err != nil {
		log.Println(err)
		return
	}
	_, err = messageSvc.CreateMessage(chat.Message{
		Text:    "HELLO" + newUser1.Name,
		OwnerID: newUser2.ID,
		RoomID:  room.ID,
	})
	if err != nil {
		log.Println(err)
		return
	}

	messages, err := messageSvc.GetMessages(nil, nil, newUser1.ID, room.ID)
	if err != nil {
		log.Println(err)
		return
	}
	for i := range messages {
		fmt.Println(messages[i].Owner.Name, messages[i].Owner.ID, "says: ", messages[i].Text)
	}
	ok, err := messageSvc.DeleteMessage(newUser1.ID, messages[1].ID)
	if err != nil {
		log.Println(ok, err)
		return
	}

	messagesAfterChanges, err := messageSvc.GetMessages(nil, nil, newUser1.ID, room.ID)
	if err != nil {
		log.Println(err)
		return
	}

	for i := range messagesAfterChanges {
		fmt.Println(messagesAfterChanges[i].Owner.Name, messagesAfterChanges[i].Owner.ID, "says: ", messagesAfterChanges[i].Text)
	}

	return
}
