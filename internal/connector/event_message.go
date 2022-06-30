package connector

const (
	ConnectToRoomEventMessageType = iota + 1
	DisconnectFromRoomEventMessageType
	NewMessageEventMessageType
)

type EventMessage struct {
	Type int8 `json:"type"`
	Data Data `json:"data"`
}

func NewEventMessage(data []byte) *EventMessage {
	return &EventMessage{}
}

type Data struct {
	*ConnectToRoomData
	*NewMessageData
}

type ConnectToRoomData struct {
	RoomID int64 `json:"room_id"`
}

type NewMessageData struct {
	Text string `json:"text"`
}
