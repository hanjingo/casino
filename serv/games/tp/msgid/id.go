package msgid

const (
	//进房
	JoinRoomReq uint32 = 20000
	JoinRoomRsp uint32 = 20001
	JoinRoomNtf uint32 = 20002

	//退房
	ExitRoomReq uint32 = 20003
	ExitRoomRsp uint32 = 20004
	ExitRoomNtf uint32 = 20005

	//坐下
	PlayerSitDownReq uint32 = 20006
	PlayerSitDownRsp uint32 = 20007
	PlayerSitDownNtf uint32 = 20008

	//玩家准备
	ReadyReq uint32 = 20009
	ReadyRsp uint32 = 20010
	ReadyNtf uint32 = 20011

	//加注
	JiaZhuReq uint32 = 20012
	JiaZhuRsp uint32 = 20013
	JiaZhuNtf uint32 = 20014

	//跟注
	GengZhuReq uint32 = 20015
	GengZhuRsp uint32 = 20016
	GengZhuNtf uint32 = 20017

	//弃牌
	QiPaiReq uint32 = 20018
	QiPaiRsp uint32 = 20019
	QiPaiNtf uint32 = 20020

	//比牌
	BiPaiReq uint32 = 20021
	BiPaiRsp uint32 = 20022
	BiPaiNtf uint32 = 20023

	//看牌
	KanPaiReq uint32 = 20024
	KanPaiRsp uint32 = 20025
	KanPaiNtf uint32 = 20026

	//不看牌
	BuKanPaiReq uint32 = 20027
	BuKanPaiRsp uint32 = 20028
	BuKanPaiNtf uint32 = 20029

	//发语音
	SendVoiceReq uint32 = 20030
	SendVoiceRsp uint32 = 20031
	SendVoiceNtf uint32 = 20032

	//发文字
	SendTextReq uint32 = 20033
	SendTextRsp uint32 = 20034
	SendTextNtf uint32 = 20035

	//发表情
	SendEmojiReq uint32 = 20036
	SendEmojiRsp uint32 = 20037
	SendEmojiNtf uint32 = 20038

	/*************通知消息****************/
	//房间关闭
	RoomCloseNtf uint32 = 21000

	//还座位
	ReturnSeatNtf uint32 = 21001

	//回合结束广播
	TurnEndNtf uint32 = 21002

	//发牌通知
	PutCardNtf uint32 = 21003

	//总结算广播
	TotalSettleNtf uint32 = 21004

	//回合结算广播
	TurnSettleNtf uint32 = 21005

	//回合开始广播
	NewTurnNtf uint32 = 21006

	//玩家掉线
	FGFPlayerOffLineNtf uint32 = 21007

	//显示牌广播
	FGFShowCardInfoNtf uint32 = 21008

	/***************计时器*******************/
	//结算倒计时
	UpdateTurnSettleTimerNtf uint32 = 22001 //更新结算倒计时
	TurnSettleCountDoneNtf   uint32 = 22002 //结算倒计时结束

	//发牌倒计时
	UpdatePutCardTimerNtf uint32 = 22003 //更新发牌倒计时
	PutCardCountDoneNtf   uint32 = 22004 //发牌倒计时结束

	//准备倒计时结束通知
	UpdateRoomReadyTimerNtf uint32 = 22005 //更新准备倒计时
	RoomReadyCountDoneNtf   uint32 = 22006 //准备倒计时结束

	//玩家操作倒计时
	UpdatePlayerOperationTimerNtf uint32 = 22007 //更新玩家操作倒计时
	PlayerOperationCountDoneNtf   uint32 = 22008 //玩家操作倒计时结束

	/***************卡住通知*******************/
	ReadyWaitingStartNtf     uint32 = 23000 //准备卡住开始
	ReadyWaitingCountDoneNtf uint32 = 23001 //准备卡住结束
)
