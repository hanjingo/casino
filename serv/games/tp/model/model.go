package model

//进房请求
type JoinRoomReq struct {
	RoomId int64 //房间id
}

//进房返回
type JoinRoomRsp struct {
	Result uint32 //结果
	Room *RoomInfo //房间信息
}

//进房广播
type JoinRoomNtf struct {
	User *PlayerInfo //用户信息
}

//退房请求
type ExitRoomReq struct {
	RoomId int64 //房间id
}

//退房返回
type ExitRoomRsp struct {
	Result uint32 //结果
}

//退房通知
type ExitRoomNtf struct {
	UserId int64 //用户id
}

//坐下请求
type SitDownReq struct {
	RoomId int64 //房间id
	SeatId uint32 //座位id
}

//坐下返回
type SitDownRsp struct {
	Result uint32 //结果
	SeatId uint32 //座位id
}

//坐下通知
type SitDownNtf struct {
	UserId int64 //用户id
	SeatId uint32 //座位id
}

//准备请求
type ReadyReq struct {
}

//准备返回
type ReadyRsp struct {
	Result uint32 //结果
}

//准备通知
type ReadyNtf struct {
	UserId int64 //用户id
}

//加注请求
type JiaZhuReq struct {
	Token string //凭证
	ChipIds []uint32 //筹码id集合
}

//加注返回
type JiaZhuRsp struct {
	Result uint32 //结果
}

//加注通知
type JiaZhuNtf struct {
	ChipIds []uint32 //筹码id集合
}

//跟注请求
type GenZhuReq struct {
	Token string //凭证
}

//跟注返回
type GenZhuRsp struct {
	Result uint32 //结果
}

//跟注通知
type GenZhuNtf struct {
	UserId int64 //用户id
}

//弃牌请求
type QiPaiReq struct {
	Token string //凭证
}

//弃牌返回
type QiPaiRsp struct {
	Result uint32 //结果
}

//弃牌通知
type QiPaiNtf struct {
	UserId int64 //用户id
}

//比牌请求
type BiPaiReq struct {
	Token string //凭证
}

//比牌返回
type BiPaiRsp struct {
	Result uint32 //结果
	WinnerId int64 //赢家
	LoserId int64 //输家
}

//弃牌通知
type BiPaiNtf struct {
	WinnerId int64 //赢家
	LoserId int64 //输家
	DoId int64 //比牌发起人
}

//看牌请求
type KanPaiReq struct {
	Token string //凭证
}

//看牌返回
type KanPaiRsp struct {
	Result uint32 //结果
}

//看牌通知
type KanPaiNtf struct {
	UserId int64 //看牌人的id
}

//房间关闭通知
type RoomCloseNtf struct {
	RoomId int64 //房间id
	Status uint32 //房间状态
}

//位置归还通知
type ReturnSeatNtf struct {
	RoomId int64 //房间id
	SeatId uint32 //座位id
	UserId int64 //用户id
}

//发牌通知
type PutCardNtf struct {
	SeatIds []int //座位id集合
}

//总结算广播
type TotalSettleNtf struct {
	RoomId int64 //房间id
	TotalTurn int //总回合数
	RoomCreateTime string //房间创建时间
	Players *PlayerTotalSettleInfo //玩家结算信息
}

//回合结算广播
type TurnSettleNtf struct {
	RoomId int64 //房间id
	TotalTurn int //回合数
	RoomCreateTime string //房间创建时间
	Players *PlayerTurnSettleInfo //玩家结算信息
}

//回合结束广播
type TurnEndNtf struct {
	RoomId int64 //房间id
	TurnNum int //当前回合
}

//回合开始广播
type TurnStartNtf struct {
	TurnNum int //当前回合
	TotalTurn int //总回合
}

//掉线广播
type OffLineNtf struct {
	UserId int64 //玩家id
}

//亮牌广播
type ShowCardNtf struct {
	SeatNum int //座位id
	Cards []int //牌集合
}

//准备倒计时刷新
type UpdateReadyTimerNtf struct {
	TimeLeft int //倒计时剩余
}

//准备倒计时结束
type ReadyTimerCountDoneNtf struct {
}

//回合结算倒计时刷新
type UpdateTurnSettleTimerNtf struct {
	TimeLeft int //倒计时剩余
}

//回合结算倒计时结束
type TurnSettleTimerCountDoneNtf struct {
}

//发牌倒计时刷新
type UpdatePutCardTimerNtf struct {
	TimeLeft int //倒计时剩余
}

//发牌倒计时结束
type PutCardTimerCountDoneNtf struct {
}

//玩家操作倒计时刷新
type UpdatePlayerOpTimerNtf struct {
	TimeLeft int //倒计时剩余
}

//玩家操作倒计时结束
type PlayerOpTimerCountDoneNtf struct {
}
