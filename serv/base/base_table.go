package base

import (
	"github.com/hanjingo/algorithm/steper"
)

type Table struct {
	Id         int           //桌子id
	BRun       bool          //桌子是否已启动
	CurrTurnId uint64        //当前回合id
	GameStatus uint32        //游戏状态
	Seats      map[int]*Seat //椅子集合
	Round      *steper.Round //计步器
}

//构造函数
func (table *Table) Table() {
	table.Id = 0
	table.BRun = false
	table.CurrTurnId = 0
	table.GameStatus = 0
	table.Seats = make(map[int]*Seat)
	table.Round = steper.NewRound()
}

//关闭桌子
func (table *Table) Close() {
	//todo
}

//创建一套桌子
func (table *Table) createSuitSeat(startIdx int, num int) {
	if num < 1 {
		return
	}
	var first *Seat
	var prev *Seat
	for i := startIdx; i < startIdx+num; i++ {
		seat := &Seat{Id: i}
		table.Seats[seat.Id] = seat
		if first == nil {
			first = seat
			prev = seat
			continue
		}
		prev.NextSeatId = seat.Id
		seat.PrevSeatId = prev.Id
		prev = seat
	}
	prev.NextSeatId = first.Id
	first.PrevSeatId = prev.Id
}

//拿到下一个座位
func (table *Table) GetNextSeat(sid int) *Seat {
	if _, ok := table.Seats[sid]; !ok {
		return nil
	}
	seat := table.Seats[sid]
	if back, ok := table.Seats[seat.NextSeatId]; ok {
		return back
	}
	return nil
}

//拿到前一个位置
func (table *Table) GetPrevSeat(sid int) *Seat {
	if _, ok := table.Seats[sid]; !ok {
		return nil
	}
	seat := table.Seats[sid]
	if back, ok := table.Seats[seat.PrevSeatId]; ok {
		return back
	}
	return nil
}

//拿座位
func (table *Table) GetSeat(sid int) *Seat {
	return table.Seats[sid]
}

//新建流程
func (table *Table) NewStep(status uint32, step steper.StepFunc) *steper.Step {
	return steper.NewStep(status, step)
}