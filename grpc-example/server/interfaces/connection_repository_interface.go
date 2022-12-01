package repository_interface

type Connection struct {
	ID string
}

type ConnectionRepository interface {
	Put(*Connection) error
	Get(ui string) (*Connection, error)
	Delete(*Connection) error
}
