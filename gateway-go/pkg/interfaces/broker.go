package interfaces

type BrokerListener interface {
	Listen() error
	Close() error
}
