package help

import (
	"github.com/hanjingo/casino/serv/base"
)

type Room struct {
	base.Room
	Players map[int64]*Player //玩家集合
	Tables  map[int]*Table    //桌子集合
	Turns   map[uint64]*Turn  //回合集合
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
