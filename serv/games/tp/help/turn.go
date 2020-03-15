package help

import (
	"github.com/hanjingo/casino/serv/base"
)

//游戏回合
type Turn struct {
	base.Turn
	Settle         *TurnSettle       //回合结算
	CurrOperSeatId int               //当前操作的位置
	Cards          *PokerSuit        //本回合的牌
	LoopNum        int               //第几轮
	LastChipNum    int               //上次筹码数
	Players        map[int64]*Player //本局参与人数
}
