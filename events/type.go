package events

type Receiver interface {
	Receive(limit int) error
}

type Processor interface {
	Process() error
}

type Sender interface {
	Send() error
}
