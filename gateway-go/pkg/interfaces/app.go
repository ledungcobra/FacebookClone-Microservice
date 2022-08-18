package interfaces

type IServer interface {
	Initialize() error
	Listen() error
	Close() error
}
