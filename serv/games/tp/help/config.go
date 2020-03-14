package help

import (
	"github.com/hanjingo/casino/serv/base"
)

func NewConfig() *TpConfig {
	back := &TpConfig{}
	back.Init()
	back.Room = &TpRoomConfig{
		TurnRoomCardMap: make(map[int]int),
		Chips:           make(map[uint32][]int),
	}
	return back
}

//赢三张配置
type TpConfig struct {
	base.Config               //基本配置
	Room        *TpRoomConfig `json:"Room"` //房间配置
}

//房间配置
type TpRoomConfig struct {
	MaxPlayerNum    int              `json:"MaxPlayerNum"`    //最大人数
	TurnRoomCardMap map[int]int      `json:"TurnRoomCardMap"` //key:总局数 value:房卡数
	MaxLimit        []int            `json:"MaxLimit"`        //封顶
	Chips           map[uint32][]int `json:"Chips"`           //筹码方案(1: 10,20,30,50,100,500 ...)
}
