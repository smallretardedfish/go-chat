package main

import (
	"github.com/smallretardedfish/go-chat/configs"
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

	owner, err := userRepo.GetUser(7)
	if err != nil {
		log.Println(err)
		return
	}

	//room, err := roomRepo.CreateRoom(room_repo.Room{
	//	Name:      "testRoom",
	//	OwnerID:   owner.ID,
	//	Owner:     owner,
	//	Type:      room_repo.PrivateRoom,
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//})

	if err != nil {
		log.Println(err)
		return
	}
	rooms, err := roomRepo.GetRooms(owner.ID)
	if err != nil {
		log.Println(err)
		return
	}
	for _, room := range rooms {
		log.Println(room.OwnerID, room.Name, room.ID)
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
	if err != nil {
		log.Println(err)
		return
	}
}
