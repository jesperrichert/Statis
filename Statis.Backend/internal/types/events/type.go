package events

type EventType int

const (
	Resolved      EventType = iota
	Update        EventType = iota
	Investigating EventType = iota
	Identified    EventType = iota
	Unkown        EventType = iota
)

var monitorName = map[EventType]string{
	Resolved:      "RESOLVED",
	Update:        "UPDATE",
	Investigating: "INVESTIGATING",
	Identified:    "IDENTIFIED",
	Unkown:        "UNKOWN",
}

func (mt EventType) ToString() string {
	return monitorName[mt]
}

func ToType(name string) EventType {
	for k, v := range monitorName {
		if v == name {
			return k
		}
	}
	return Unkown
}
