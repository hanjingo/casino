package lib

import "time"

type TurnBase struct {
	turnNum    int       //所属回合
	createTime time.Time //创建时间
	endTime    time.Time //结束时间
}

func (base *TurnBase) GetTurnNum() int {
	return base.turnNum
}

func (base *TurnBase) SetTurnNum(n int) {
	base.turnNum = n
}

func (base *TurnBase) GetCreateTime() time.Time {
	return base.createTime
}

func (base *TurnBase) SetCreateTime(t time.Time) {
	base.createTime = t
}

func (base *TurnBase) GetEndTime() time.Time {
	return base.endTime
}

func (base *TurnBase) SetEndTime(t time.Time) {
	base.endTime = t
}
