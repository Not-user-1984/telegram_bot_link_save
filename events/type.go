package events

type Fatcher interface {
	Fatch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
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