package events

type Fatcher interface {
	Fatch(limit int) ([]Event, error)
}

type Processor interface {
	Processor(e Event) error
}

type Type int

const(
	Unknow Type = iota
	message
)

type Event struct {
	Type Type
	Text string
}