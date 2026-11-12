package monitors

import (
	"encoding/json"

	"jespersen.zip.statis/internal/dto"
)

type MonitorType int

type DummyMonitor struct {
}

type HttpMonitor struct {
	Method string `json:"method"`
	Url    string `json:"url"`
}

type TcpMonitor struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type PingMonitor struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

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

func (mt MonitorType) IsTypeDtoValid(monitor dto.MonitorDto[any]) bool {
	var bytes []byte
	bytes, _ = json.Marshal(monitor)
	switch mt {
	case Group:
		{
			err := json.Unmarshal(bytes, &dto.MonitorDto[DummyMonitor]{})
			if err == nil {
				return true
			}
		}
	case Tcp:
		{
			err := json.Unmarshal(bytes, &dto.MonitorDto[TcpMonitor]{})
			if err == nil {
				return true
			}
		}
	case Http:
		{
			err := json.Unmarshal(bytes, &dto.MonitorDto[HttpMonitor]{})
			if err == nil {
				return true
			}
		}
	case Ping:
		{
			err := json.Unmarshal(bytes, &dto.MonitorDto[PingMonitor]{})
			if err == nil {
				return true
			}
		}
	default:
		return false
	}

	return false
}

func ToType(name string) MonitorType {
	for k, v := range monitorTypeName {
		if v == name {
			return k
		}
	}
	return Unkown
}
