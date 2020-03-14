package model

func Init() {
	var reg = h.GetDefaultModelMap().Reg

	//新到连接
	reg(com.NewConnReqId, com.NewConnReq{})
	reg(com.NewConnRspId, com.NewConnRsp{})

	//进房
	reg(msgid.JoinRoomReq, JoinRoomReq{})
	reg(msgid.JoinRoomRsp, JoinRoomRsp{})
	reg(msgid.JoinRoomNtf, JoinRoomNtf{})

	//退房
	reg(msgid.ExitRoomReq, ExitRoomReq{})
	reg(msgid.ExitRoomRsp, ExitRoomRsp{})
	reg(msgid.ExitRoomNtf, ExitRoomNtf{})

	//坐下
	reg(msgid.PlayerSitDownReq, PlayerSitDownReq{})
	reg(msgid.PlayerSitDownRsp, PlayerSitDownRsp{})
	reg(msgid.PlayerSitDownNtf, PlayerSitDownNtf{})

	//玩家准备
	reg(msgid.ReadyReq, ReadyReq{})
	reg(msgid.ReadyRsp, ReadyRsp{})
	reg(msgid.ReadyNtf, ReadyNtf{})

	//加注
	reg(msgid.JiaZhuReq, JiaZhuReq{})
	reg(msgid.JiaZhuRsp, JiaZhuRsp{})
	reg(msgid.JiaZhuNtf, JiaZhuNtf{})

	//跟注
	reg(msgid.GengZhuReq, GengZhuReq{})
	reg(msgid.GengZhuRsp, GengZhuRsp{})
	reg(msgid.GengZhuNtf, GengZhuNtf{})

	//弃牌
	reg(msgid.QiPaiReq, QiPaiReq{})
	reg(msgid.QiPaiRsp, QiPaiRsp{})
	reg(msgid.QiPaiNtf, QiPaiNtf{})

	//比牌
	reg(msgid.BiPaiReq, BiPaiReq{})
	reg(msgid.BiPaiRsp, BiPaiRsp{})
	reg(msgid.BiPaiNtf, BiPaiNtf{})

	//看牌
	reg(msgid.KanPaiReq, KanPaiReq{})
	reg(msgid.KanPaiRsp, KanPaiRsp{})
	reg(msgid.KanPaiNtf, KanPaiNtf{})

	//不看牌
	reg(msgid.BuKanPaiReq, BuKanPaiReq{})
	reg(msgid.BuKanPaiRsp, BuKanPaiRsp{})
	reg(msgid.BuKanPaiNtf, BuKanPaiNtf{})

	//发语音
	reg(msgid.SendVoiceReq, SendVoiceReq{})
	reg(msgid.SendVoiceRsp, SendVoiceRsp{})
	reg(msgid.SendVoiceNtf, SendVoiceNtf{})

	//发文字
	reg(msgid.SendTextReq, SendTextReq{})
	reg(msgid.SendTextRsp, SendTextRsp{})
	reg(msgid.SendTextNtf, SendTextNtf{})

	//发表情
	reg(msgid.SendEmojiReq, SendEmojiReq{})
	reg(msgid.SendEmojiRsp, SendEmojiRsp{})
	reg(msgid.SendEmojiNtf, SendEmojiNtf{})

}