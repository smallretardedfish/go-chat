package chat

type RoomService interface {
	GetRoom(userID, roomID int64) (*Room, error)
	GetRooms(limit, offset, userID int64) ([]Room, error)
	GetUsers(roomID, userID int64) ([]User, error)
	CreateRoom(room Room, userIDs []int64) (*Room, error)
	UpdateRoom(userID int64, room Room) (*Room, error)
	DeleteRoom(userID, roomID int64) (*Room, error)
	AddUserToRoom(userID, roomID int64) (bool, error)
	DeleteUserFromRoom(userID, roomID int64) (bool, error)
}
