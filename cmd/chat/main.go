package main

import (
	"fmt"
	"github.com/smallretardedfish/go-chat/configs"
	"github.com/smallretardedfish/go-chat/internal/repositories/message_repo"
	"github.com/smallretardedfish/go-chat/internal/repositories/room_repo"
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
	roomRepo := room_repo.NewRoomRepo(db)
	messageRepo := message_repo.NewMessageRepo(db)
	rooms, err := roomRepo.GetRooms(7)
	if err != nil {
		log.Println(err)
		return
	}
	for _, room := range rooms {
		fmt.Println(room.ID, room.Name, room.Owner.Name, room.Owner.ID)
	}
	//msg, err := messageRepo.GetMessage(10)
	//if err != nil {
	//	log.Println(err)
	//}

	m, err := messageRepo.CreateMessage(message_repo.Message{
		Text:    "testing owner populating 2",
		OwnerID: 7,
		RoomID:  10,
	})
	mes, err := messageRepo.GetMessage(m.ID)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(mes.Owner.Name)
	//
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
