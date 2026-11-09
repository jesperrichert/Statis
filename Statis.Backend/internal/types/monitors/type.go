package monitors

type MonitorType int

const (
	Group  MonitorType = iota
	Tcp    MonitorType = iota
	Http   MonitorType = iota
	Ping   MonitorType = iota
	Unkown MonitorType = iota
)

var monitorTypeName = map[MonitorType]string{
	Group:  "GROUP",
	Tcp:    "TCP",
	Http:   "HTTP",
	Ping:   "PING",
	Unkown: "UNKOWN",
}

func (mt MonitorType) ToString() string {
	return monitorTypeName[mt]
}

func ToType(name string) MonitorType {
	for k, v := range monitorTypeName {
		if v == name {
			return k
		}
	}
	return Unkown
}
