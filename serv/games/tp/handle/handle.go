package handle

import (
	"github.com/hanjingo/casino/serv/games/tp/model"
	"github.com/hanjingo/casino/serv/games/tp/msgid"
	"github.com/hanjingo/casino/serv/lib"
)

//进房
func OnJoinRoomReq(session *lib.Session, req *model.JoinRoomReq) {
	rsp := &model.JoinRoomRsp{Result: 0}
	defer session.Gate.Rsp(msgid.JoinRoomRsp, rsp)
}

//退房
func OnExitRoomReq(session *lib.Session, req *model.ExitRoomReq) {
	rsp := &model.ExitRoomRsp{Result: 0}
	defer session.Gate.Rsp(msgid.ExitRoomRsp, rsp)
}

//坐下
func OnSitDownReq(session *lib.Session, req *model.SitDownReq) {
	rsp := &model.SitDownRsp{Result: 0}
	defer session.Gate.Rsp(msgid.SitDownRsp, rsp)
}

//准备
func OnReadyReq(session *lib.Session, req *model.ReadyReq) {
	rsp := &model.ReadyRsp{Result: 0}
	defer session.Gate.Rsp(msgid.ReadyRsp, rsp)
}

//加注
func OnJiaZhuReq(session *lib.Session, req *model.JiaZhuReq) {
	rsp := &model.JiaZhuRsp{Result: 0}
	defer session.Gate.Rsp(msgid.JiaZhuRsp, rsp)
}

//跟注
func OnGenZhuReq(session *lib.Session, req *model.GenZhuReq) {
	rsp := &model.GenZhuRsp{Result: 0}
	defer session.Gate.Rsp(msgid.GenZhuRsp, rsp)
}

//弃牌
func OnQiPaiReq(session *lib.Session, req *model.QiPaiReq) {
	rsp := &model.QiPaiRsp{Result: 0}
	defer session.Gate.Rsp(msgid.QiPaiRsp, rsp)
}

//比牌
func OnBiPaiReq(session *lib.Session, req *model.BiPaiReq) {
	rsp := &model.BiPaiRsp{Result: 0}
	defer session.Gate.Rsp(msgid.BiPaiRsp, rsp)
}

//看牌
func OnKanPaiReq(session *lib.Session, req *model.KanPaiReq) {
	rsp := &model.KanPaiRsp{Result: 0}
	defer session.Gate.Rsp(msgid.KanPaiRsp, rsp)
}
