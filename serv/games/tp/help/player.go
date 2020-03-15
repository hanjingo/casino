package help

import (
	"github.com/hanjingo/casino/serv/base"
)

type Player struct {
	base.Player
	CurrSeat  *base.Seat //玩家位置
	HandCards []*Poker   //玩家手牌
}
