package msgid

const (
	JoinRoomReq uint32 = 1001 //进房请求

	JoinRoomRsp uint32 = 1002 //进房返回

	JoinRoomNtf uint32 = 1003 //进房广播

	ExitRoomReq uint32 = 1004 //退房请求

	ExitRoomRsp uint32 = 1005 //退房返回

	ExitRoomNtf uint32 = 1006 //退房通知

	SitDownReq uint32 = 1007 //坐下请求

	SitDownRsp uint32 = 1008 //坐下返回

	SitDownNtf uint32 = 1009 //坐下通知

	ReadyReq uint32 = 1010 //准备请求

	ReadyRsp uint32 = 1011 //准备返回

	ReadyNtf uint32 = 1012 //准备通知

	JiaZhuReq uint32 = 1013 //加注请求

	JiaZhuRsp uint32 = 1014 //加注返回

	JiaZhuNtf uint32 = 1015 //加注通知

	GenZhuReq uint32 = 1016 //跟注请求

	GenZhuRsp uint32 = 1017 //跟注返回

	GenZhuNtf uint32 = 1018 //跟注通知

	QiPaiReq uint32 = 1019 //弃牌请求

	QiPaiRsp uint32 = 1020 //弃牌返回

	QiPaiNtf uint32 = 1021 //弃牌通知

	BiPaiReq uint32 = 1022 //比牌请求

	BiPaiRsp uint32 = 1023 //比牌返回

	BiPaiNtf uint32 = 1024 //弃牌通知

	KanPaiReq uint32 = 1025 //看牌请求

	KanPaiRsp uint32 = 1026 //看牌返回

	KanPaiNtf uint32 = 1027 //看牌通知

	RoomCloseNtf uint32 = 1100 //房间关闭通知

	ReturnSeatNtf uint32 = 1101 //位置归还通知

	PutCardNtf uint32 = 1102 //发牌通知

	TotalSettleNtf uint32 = 1103 //总结算广播

	TurnSettleNtf uint32 = 1104 //回合结算广播

	TurnEndNtf uint32 = 1105 //回合结束广播

	TurnStartNtf uint32 = 1106 //回合开始广播

	OffLineNtf uint32 = 1107 //掉线广播

	ShowCardNtf uint32 = 1108 //亮牌广播

	UpdateReadyTimerNtf uint32 = 1200 //准备倒计时刷新

	ReadyTimerCountDoneNtf uint32 = 1201 //准备倒计时结束

	UpdateTurnSettleTimerNtf uint32 = 1202 //回合结算倒计时刷新

	TurnSettleTimerCountDoneNtf uint32 = 1203 //回合结算倒计时结束

	UpdatePutCardTimerNtf uint32 = 1204 //发牌倒计时刷新

	PutCardTimerCountDoneNtf uint32 = 1205 //发牌倒计时结束

	UpdatePlayerOpTimerNtf uint32 = 1206 //玩家操作倒计时刷新

	PlayerOpTimerCountDoneNtf uint32 = 1207 //玩家操作倒计时结束
)
