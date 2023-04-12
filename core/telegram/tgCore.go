package telegram

import (
	"PersonalPlanner/events"
	"log"
	"time"
)

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
		newEvents, err := c.receiver.Receive(c.batchSize)
		if err != nil {
			log.Printf("Error in tgCore recieve: %s", err)

			continue
		}

		if len(newEvents) == 0 {
			time.Sleep(2 * time.Second)

			continue
		}

		if err = c.handleEvents(newEvents); err != nil {
			log.Print(err)

			continue
		}
	}
}

// TODO handle errors and parallels
func (c TgCore) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err)

			continue
		}
	}

	return nil
}
