package help

import (
	"context"
	"math/rand"
	"time"

	"github.com/hanjingo/casino/serv/base"
	"github.com/hanjingo/casino/serv/errid"
	"github.com/hanjingo/casino/serv/games/tp/model"
	"github.com/hanjingo/util"
)

type Table struct {
	base.Table
	LastWinnerId int64            //上局赢家
	TotalSettle  *TotalSettle     //总结算
	Turns        map[uint64]*Turn //回合集合
	CurrTurnId   uint64           //当前回合id
	CurrRoomId   int64            //这张桌子所在的房间
}

func newTable(id int) *Table {
	back := new(Table)
	back.Table.Table()
	back.LastWinnerId = 0
	back.TotalSettle = newTotalSettle()
	return back
}

//建立新回合
func (table *Table) NewTurn() uint32 {
	//todo
	return errid.OK.Id()
}

//拿到当前回合
func (table *Table) GetCurrTurn() *Turn {
	if turn, ok := table.Turns[table.CurrTurnId]; ok {
		return turn
	}
	return nil
}

//获得这张桌子所在的房间
func (table *Table) GetCurrRoom() *Room {
	back := GameApp.RoomM.GetRoom(table.CurrRoomId)
	if back != nil {
		return back.(*Room)
	}
	return nil
}

/*
1.准备
2.开始新回合
3.发牌
4.轮转
5.结算
6.开始新回合
*/
//跑起来
func (table *Table) Run() {
	if table.BRun {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("房间运行失败,错误:%v", util.GetRunTimeStack())
		}
	}()
	table.BRun = true
	table.GameStatus = GAMING_STATUS_START
	table.Round.PushBack(table.NewStep(GAMING_STATUS_READYING, table.ready)) //准备
	go table.Round.Run()
}

//准备
func (table *Table) ready(ctx context.Context) uint32 {
	log.Debug("Ready>>")
	defer func() {
		table.BroadCast(&model.ReadyTimerCountDoneNtf{})
		//开始回合
		if err := table.NewTurn(); err != errid.OK.Id() {
			log.Error("桌子id:%v新建回合失败!!!", table.Id)
			table.Close()
			return
		}
	}()
	table.GameStatus = GAMING_STATUS_READYING
	tm := time.NewTimer(time.Duration(1 * time.Second))
	readyTimeLeft := Config.Room.ReadyTimeOut
	for {
		select {
		case <-ctx.Done():
			return GAMING_STATUS_READY_DONE
		case <-tm.C:
			if readyTimeLeft <= 0 {
				return GAMING_STATUS_READY_DONE
			}
			ntf := &model.UpdateReadyTimerNtf{TimeLeft: int32(readyTimeLeft)}
			table.BroadCast(ntf)
			readyTimeLeft--
			tm.Reset(time.Duration(1 * time.Second))
		}
	}
	return GAMING_STATUS_READY_DONE
}

//发牌
func (table *Table) putCard(ctx context.Context) uint32 {
	log.Debug("putCard>>")
	defer func() {
		table.BroadCast(&model.PutCardTimerCountDoneNtf{})

		//找到开始的玩家

	}()
	//发牌

	return GAMING_STATUS_PUT_CARD_DONE
}

//玩家操作
func (table *Table) loop(ctx context.Context) uint32 {
	//todo
	return GAMING_STATUS_LOOP_DONE
}

//结算
func (table *Table) settle(ctx context.Context) uint32 {
	//todo
	return GAMING_STATUS_TURN_SETTLE_DONE
}

//执行发牌
func (table *Table) doPutCard() {
	turn := table.GetCurrTurn()
	if turn == nil {
		log.Error("发牌时当前回合不存在")
		return
	}
	var players []*Player
	//先找到开始发牌的人 (如果上把有赢家，就从他开始发)
	if p, ok := turn.Players[table.LastWinnerId]; ok {
		players = append(players, p)
	} else {
		//随机找人
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		n := r.Intn(len(turn.Players))
		for _, p := range turn.Players {
			if n == 0 {
				players = append(players, p)
			}
			n--
		}
	}
	if len(players) < 1 {
		return
	}
	firstP := players[0]
	if firstP.CurrSeat == nil {

	}
	turn.CurrOperSeatId = firstP.CurrSeat.Id
	for {
		nextSeat := table.GetNextSeat(turn.CurrOperSeatId)
		temp := turn.Players[nextSeat.PlayerId]
		if temp == nil {
			continue
		}
		if temp.GetId() == firstP.GetId() {
			break
		}
		players = append(players, temp)
	}

	//开始发牌
	turn.Cards = NewPokerSuit()
	turn.Cards.Init()
	turn.Cards.Shuffle() //洗牌

	var seats []int
	for _, p := range players {
		p.HandCards = turn.Cards.PopCard(3)
		seats = append(seats, p.CurrSeat.Id)
	}
	table.BroadCast(&model.PutCardNtf{SeatIds: seats})
}

//做坏事 嘿嘿
func (table *Table) doBad() {
	room := table.GetCurrRoom()
	if room == nil {
		log.Error("发牌时找不到房子")
		return
	}
}
