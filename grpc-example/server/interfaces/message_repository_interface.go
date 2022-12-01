package repository_interface

type Message struct {
	ID   string
	Rh   string
	X    int
	Y    int
	From string
}

type MessageRepository interface {
	Put(*Message) error
	Get() ([]*Message, error)
	Delete(messageIDs []*Message) error
}
