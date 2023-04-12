package telegram

import "PersonalPlanner/events"

type TgCore struct {
	receiver  events.Receiver
	processor events.Processor
	sender    events.Sender
	batchSize int
}

func New(receiver events.Receiver, processor events.Processor, sender events.Sender, batchSize int) TgCore {
	return TgCore{
		receiver:  receiver,
		processor: processor,
		sender:    sender,
		batchSize: batchSize,
	}
}

func (c TgCore) Start() error {
	for {

	}
}
