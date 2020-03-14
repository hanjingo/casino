package base

import "time"

type Turn struct {
	TurnNum    int       //所属回合
	CreateTime time.Time //创建时间
	EndTime    time.Time //结束时间
}

//回合开始
func (turn *Turn) Start() {
	turn.CreateTime = time.Now()
}

//回合结束
func (turn *Turn) End() {
	turn.EndTime = time.Now()
}

//回合结算
func (turn *Turn) Settle() {
	//
}
