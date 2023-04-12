package core

import "PersonalPlanner/events"

type Core interface {
	events.Fetcher
	events.Processor
	Start() error
}
