package main

import (
	"fmt"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/repositories/message_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/room_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/user_repo"
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

	var names = []string{"Dude", "Joe"}
	var users []*user_repo.User

	for _, name := range names {
		user, err := userRepo.CreateUser(user_repo.User{
			Name: name,
		})
		if err != nil {
			log.Println("error while creating user:", err)
			return
		}
		users = append(users, user)
		_, err = userRepo.CreateUserCredentials(user_repo.UserCredentials{
			ID:       user.ID,
			Password: name + "Password",
		})

		if err != nil {
			log.Println("error while creating user credentials:", err)
		}
	}
	room, _ := roomRepo.CreateRoom(room_repo.Room{
		Name:    "testRoom",
		OwnerID: users[0].ID,
		Type:    room_repo.PublicRoom,
	})

	updatedRoom := room
	updatedRoom.RoomUsers = append(updatedRoom.RoomUsers, room_repo.RoomUser{
		RoomID: room.ID,
		UserID: users[1].ID,
		Status: room_repo.RoomUserCreated,
	})
	roomRepo.UpdateRoom(room.OwnerID, *updatedRoom)

	//msg, err := messageRepo.GetMessage(10)
	//if err != nil {
	//	log.Println(err)
	//}

	m, err := messageRepo.CreateMessage(message_repo.Message{
		Text:    "testing message creation 2",
		OwnerID: users[1].ID,
		RoomID:  room.ID,
	})
	mes, err := messageRepo.GetMessage(m.ID)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(mes.Owner.Name)

	//msg, err = messageRepo.UpdateMessage(message_repo.Message{
	//	ID:        msg.ID,
	//	Text:      "edited",
	//	OwnerID:   msg.OwnerID,
	//	Owner:     msg.Owner,
	//	RoomID:    msg.RoomID,
	//	CreatedAt: msg.CreatedAt,
	//	UpdatedAt: time.Now(),
	//})
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//msgs, err := messageRepo.AllMessagesInRoom(10)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//for _, m := range msgs {
	//	log.Println(m.Owner.Name, m.OwnerID, ":", m.Text)
	//}

	//room, err := roomRepo.CreateRoom(room_repo.Room{
	//	Name:      "testRoom",
	//	OwnerID:   owner.ID,
	//	Owner:     owner,
	//	Type:      room_repo.PrivateRoom,
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//})

}

//var roomUsers []room_repo.RoomUser
//var users []user_repo.User

//db.Model(user_repo.User{}).Find(&users) // all users including owner
//
//for _, user := range users {
//	roomUsers = append(roomUsers, room_repo.RoomUser{
//		RoomID: room.ID,
//		UserID: user.ID,
//		Status: room_repo.RoomUserCreated,
//	})
//}
//_, err = roomRepo.UpdateRoom(room.OwnerID, room_repo.Room{
//	ID:      room.ID,
//	Name:    room.Name,
//	OwnerID: room.OwnerID,
//	Type:    room.Type,
//	//	RoomUsers: roomUsers,
//	CreatedAt: room.CreatedAt,
//	UpdatedAt: room.UpdatedAt,
//})
