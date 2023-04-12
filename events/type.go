package events

type Fetcher interface {
	Fetch(limit int) error
}

type Processor interface {
	Process() error
}
