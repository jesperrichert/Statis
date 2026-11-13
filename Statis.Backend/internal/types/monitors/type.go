package monitors

import (
	"encoding/json"

	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/util"
)

type MonitorType int

type DummyMonitor struct{}

type HttpMonitor struct {
	Method string `json:"method"`
	Url    string `json:"url"`
	Body   string `json:"body"`
	Status int    `json:"status"`
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
		return util.CanUnmarshalAs[dto.MonitorDto[DummyMonitor]](bytes)
	case Tcp:
		return util.CanUnmarshalAs[dto.MonitorDto[TcpMonitor]](bytes)
	case Http:
		return util.CanUnmarshalAs[dto.MonitorDto[HttpMonitor]](bytes)
	case Ping:
		return util.CanUnmarshalAs[dto.MonitorDto[PingMonitor]](bytes)
	default:
		return false
	}
}

func ToType(name string) MonitorType {
	for k, v := range monitorTypeName {
		if v == name {
			return k
		}
	}
	return Unkown
}
