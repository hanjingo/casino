package handle

import (
	"github.com/hanjingo/casino/serv/games/tp/model"
	"github.com/hanjingo/casino/serv/games/tp/msgid"
	v1 "github.com/hanjingo/gate/cli/golang/v1"
)

//进房
func OnJoinRoomReq(gate *v1.GateCli, agentId uint64, req *model.JoinRoomReq) {
	rsp := &model.JoinRoomRsp{Result: 0}
	//todo
	gate.Rsp(msgid.JoinRoomRsp, rsp, agentId)
}

//退房
func OnExitRoomReq(gate *v1.GateCli, agentId uint64, req *model.ExitRoomReq) {
	rsp := &model.ExitRoomRsp{Result: 0}
	//todo
	gate.Rsp(msgid.ExitRoomRsp, rsp, agentId)
}

//坐下
func OnSitDownReq(gate *v1.GateCli, agentId uint64, req *model.SitDownReq) {
	rsp := &model.SitDownRsp{Result: 0}
	//todo
	gate.Rsp(msgid.SitDownRsp, rsp, agentId)
}

//准备
func OnReadyReq(gate *v1.GateCli, agentId uint64, req *model.ReadyReq) {
	rsp := &model.ReadyRsp{Result: 0}
	//todo
	gate.Rsp(msgid.ReadyRsp, rsp, agentId)
}

//加注
func OnJiaZhuReq(gate *v1.GateCli, agentId uint64, req *model.JiaZhuReq) {
	rsp := &model.JiaZhuRsp{Result: 0}
	//todo
	gate.Rsp(msgid.JiaZhuRsp, rsp, agentId)
}

//跟注
func OnGenZhuReq(gate *v1.GateCli, agentId uint64, req *model.GenZhuReq) {
	rsp := &model.GenZhuRsp{Result: 0}
	//todo
	gate.Rsp(msgid.GenZhuRsp, rsp, agentId)
}

//弃牌
func OnQiPaiReq(gate *v1.GateCli, agentId uint64, req *model.QiPaiReq) {
	rsp := &model.QiPaiRsp{Result: 0}
	//todo
	gate.Rsp(msgid.QiPaiRsp, rsp, agentId)
}

//比牌
func OnBiPaiReq(gate *v1.GateCli, agentId uint64, req *model.BiPaiReq) {
	rsp := &model.BiPaiRsp{Result: 0}
	//todo
	gate.Rsp(msgid.BiPaiRsp, rsp, agentId)
}

//看牌
func OnKanPaiReq(gate *v1.GateCli, agentId uint64, req *model.KanPaiReq) {
	rsp := &model.KanPaiRsp{Result: 0}
	//todo
	gate.Rsp(msgid.KanPaiRsp, rsp, agentId)
}
