package base

import (
	l "github.com/hanjingo/logger"
	"github.com/hanjingo/network"
)

//基本信息配置
type Config struct {
	Version   string                 `json:"Version"`   //版本
	Id        uint64                 `json:"Id"`        //id
	GameType  uint32                 `json:"GameType"`  //游戏类型
	Port      uint32                 `json:"Port"`      //端口
	ScanDur   int                    `json:"ScanDur"`   //扫描间隔
	GateAddr  string                 `json:"GateAddr"`  //网关地址
	GateToken string                 `json:"GateToken"` //网关token
	Log       *l.LogConfig           `json:"Log"`       //日志配置
	Net       *network.SessionConfig `json:"Net"`       //网络配置
}

func (conf *Config) Init() {
	conf.Log = l.NewConfig()
	conf.Net = &network.SessionConfig{}
}
