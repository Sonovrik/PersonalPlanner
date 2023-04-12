package core

import "PersonalPlanner/events"

type Core interface {
	events.Receiver
	events.Processor
	events.Sender
	Start() error
}
