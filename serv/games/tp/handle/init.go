package handle

import (
	"github.com/hanjingo/casino/serv/games/tp/model"
	"github.com/hanjingo/casino/serv/games/tp/msgid"
	"github.com/hanjingo/casino/serv/lib"
)

func Init(handler *lib.Handler) {
	reg := handler.Reg

	//进房
	reg(msgid.JoinRoomReq, OnJoinRoomReq, model.JoinRoomReq{})
	//退房
	reg(msgid.ExitRoomReq, OnExitRoomReq, model.ExitRoomReq{})
	//坐下
	reg(msgid.SitDownReq, OnSitDownReq, model.SitDownReq{})
	//准备
	reg(msgid.ReadyReq, OnReadyReq, model.ReadyReq{})
	//加注
	reg(msgid.JiaZhuReq, OnJiaZhuReq, model.JiaZhuReq{})
	//跟注
	reg(msgid.GenZhuReq, OnGenZhuReq, model.GenZhuReq{})
	//弃牌
	reg(msgid.QiPaiReq, OnQiPaiReq, model.QiPaiReq{})
	//比牌
	reg(msgid.BiPaiReq, OnBiPaiReq, model.BiPaiReq{})
	//看牌
	reg(msgid.KanPaiReq, OnKanPaiReq, model.KanPaiReq{})
}
