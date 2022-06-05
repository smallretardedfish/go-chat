package message_repo

type MessageFilter struct {
	Search *string
	Limit  *int64
	Offset *int64
}
