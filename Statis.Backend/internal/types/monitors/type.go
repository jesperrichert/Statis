package monitors

type MonitorType int

const (
	Group  MonitorType = iota
	Tcp    MonitorType = iota
	Http   MonitorType = iota
	Ping   MonitorType = iota
	Unkown MonitorType = iota
)

var monitorName = map[MonitorType]string{
	Group:  "GROUP",
	Tcp:    "TCP",
	Http:   "HTTP",
	Ping:   "PING",
	Unkown: "UNKOWN",
}

func (mt MonitorType) ToString() string {
	return monitorName[mt]
}

func ToType(name string) MonitorType {
	for k, v := range monitorName {
		if v == name {
			return k
		}
	}
	return Unkown
}
