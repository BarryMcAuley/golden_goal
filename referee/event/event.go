package event

const (
	EventNewMatch = 1 << iota
)

type Event struct {
	EventType     int
	EventSource   string
	EventMessage  string
	EventTeamHome string
	EventTeamAway string
}

func NewMatchEvent(home string, away string) *Event {
	return &Event{
		EventType:     EventNewMatch,
		EventTeamHome: home,
		EventTeamAway: away,
	}
}
