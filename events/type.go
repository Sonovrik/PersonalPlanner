package events

type Receiver interface {
	Receive(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Sender interface {
	Send() error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
	Meta interface{}
}
