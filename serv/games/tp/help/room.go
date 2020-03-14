package help

import (
	"github.com/hanjingo/algorithm/steper"
	"github.com/hanjingo/casino/serv/base"
)

type Room struct {
	base.Room
	Round      *steper.Round //计步器
	GameStatus uint32        //游戏状态
	
}

//初始化
func (room *Room) Init() {

}

//关闭
func (room *Room) Close() {

}

//结算
func (room *Room) Settle() {

}
