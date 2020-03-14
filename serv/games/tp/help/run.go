package help

import (
	"context"

	"github.com/hanjingo/algorithm/steper"
	"github.com/hanjingo/util"
)

//跑起来
func (room *Room) Run() {
	if room.BRun {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("房间运行失败,错误:%v", util.GetRunTimeStack())
		}
	}()
	room.BRun = true
	room.GameStatus = GAMING_STATUS_START
	room.Round.PushBack(steper.NewStep(GAMING_STATUS_READYING, room.ready)) //准备
	go room.Round.Run()
}

//准备
func (room *Room) ready(ctx context.Context) uint32 {
	//todo
	return GAMING_STATUS_READY_DONE
}

//发牌
func (room *Room) putCard(ctx context.Context) uint32 {
	//todo
	return GAMING_STATUS_PUT_CARD_DONE
}

//玩家操作
func (room *Room) loop(ctx context.Context) uint32 {
	//todo
	return GAMING_STATUS_LOOP_DONE
}

//结算
func (room *Room) settle(ctx context.Context) uint32 {
	//todo
	return GAMING_STATUS_TURN_SETTLE_DONE
}
