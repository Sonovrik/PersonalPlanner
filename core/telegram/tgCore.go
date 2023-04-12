package telegram

import "PersonalPlanner/events"

type TgCore struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) TgCore {
	return TgCore{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c TgCore) Start() error {
	for {

	}
}
