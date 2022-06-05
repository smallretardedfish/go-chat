package chat

type Message struct {
	ID    int64
	Text  string
	Owner User
}
