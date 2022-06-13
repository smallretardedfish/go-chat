package room_repo

type RoomFilter struct {
	Search *string
	Limit  *int64
	Offset *int64
}
