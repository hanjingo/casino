package help

import (
	l "github.com/hanjingo/logger"
)

type TpConfig struct {
	Version string        `json:"Version"` //版本
	Base    *BaseConfig   `json:"Base"`    //基本配置
	Room    *TpRoomConfig `json:"Room"`    //房间配置
	Log     *LogConfig    `json:"Log"`     //日志配置
}

type BaseConfig struct {
	Id       uint32 `json:"Id"`       //id
	GameType uint32 `json:"GameType"` //游戏类型
	Port     uint32 `json:"Port"`     //端口
	ScanDur  int    `json:"ScanDur"`  //扫描间隔
}

type TpRoomConfig struct {
	MaxPlayerNum    int              `json:"MaxPlayerNum"`    //最大人数
	TurnRoomCardMap map[int]int      `json:"TurnRoomCardMap"` //key:总局数 value:房卡数
	MaxLimit        []int            `json:"MaxLimit"`        //封顶
	Chips           map[uint32][]int `json:"Chips"`           //筹码方案(1: 10,20,30,50,100,500 ...)
}

type LogConfig struct {
	DoPrintConsole bool                   `json:"DoPrintConsole"` //是否打印终端
	DoPrintFile    bool                   `json:"DoPrintFile"`    //是否打印文件
	DoPrintNet     bool                   `json:"DoPrintNet"`     //是否打印网路
	Console        *l.ConsoleWriterConfig `json:"Console"`        //终端打印机配置
	File           *l.FileWriterConfig    `json:"File"`           //文件打印机配置
	Net            *l.NetWriterConfig     `json:"Net"`            //网络打印机配置
}
