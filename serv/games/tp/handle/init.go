package handle

import (
	"github.com/hanjingo/casino/serv/games/tp/model"
	"github.com/hanjingo/casino/serv/games/tp/msgid"
	v1 "github.com/hanjingo/gate/cli/golang/v1"
)

var Gate *v1.GateCli

func Init() {
	//进房
	Gate.RegHandler(msgid.JoinRoomReq, OnJoinRoomReq, model.JoinRoomReq{})
	//退房
	Gate.RegHandler(msgid.ExitRoomReq, OnExitRoomReq, model.ExitRoomReq{})
	//坐下
	Gate.RegHandler(msgid.SitDownReq, OnSitDownReq, model.SitDownReq{})
	//准备
	Gate.RegHandler(msgid.ReadyReq, OnReadyReq, model.ReadyReq{})
	//加注
	Gate.RegHandler(msgid.JiaZhuReq, OnJiaZhuReq, model.JiaZhuReq{})
	//跟注
	Gate.RegHandler(msgid.GenZhuReq, OnGenZhuReq, model.GenZhuReq{})
	//弃牌
	Gate.RegHandler(msgid.QiPaiReq, OnQiPaiReq, model.QiPaiReq{})
	//比牌
	Gate.RegHandler(msgid.BiPaiReq, OnBiPaiReq, model.BiPaiReq{})
	//看牌
	Gate.RegHandler(msgid.KanPaiReq, OnKanPaiReq, model.KanPaiReq{})
}
