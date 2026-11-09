package monitors

type MonitorState int

const (
	Online   MonitorState = iota
	Offline  MonitorState = iota
	Degraded MonitorState = iota
	Effected MonitorState = iota
)

var monitorStateName = map[MonitorState]string{
	Degraded: "degraded",
	Offline:  "offline",
	Online:   "online",
	Effected: "effected",
}

func (mt MonitorState) ToString() string {
	return monitorStateName[mt]
}

func ToState(name string) MonitorState {
	for k, v := range monitorStateName {
		if v == name {
			return k
		}
	}
	return Offline
}
